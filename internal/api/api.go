package api

import (
	"context"
	"github.com/ocp-docs-api/internal/flusher"
	"github.com/ocp-docs-api/internal/metrics"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/producer"
	"github.com/ocp-docs-api/internal/repo"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpDocsApiServer
	repo    repo.Repo
	flusher flusher.Flusher
	prod    producer.Producer
}

func toMessage(doc document.Document) *desc.Doc {
	return &desc.Doc{
		Id:         doc.Id,
		Name:       doc.Name,
		Link:       doc.Link,
		SourceLink: doc.SourceLink,
	}
}

func fromMessageDoc(doc *desc.Doc) document.Document {
	return document.Document{
		Id:         doc.Id,
		Name:       doc.Name,
		Link:       doc.Link,
		SourceLink: doc.SourceLink,
	}
}

func fromMessageNewDoc(doc *desc.NewDoc, id uint64) document.Document {
	return document.Document{
		Id:         id,
		Name:       doc.Name,
		Link:       doc.Link,
		SourceLink: doc.SourceLink,
	}
}

func NewDocsApi(repo repo.Repo, flusher flusher.Flusher, prod producer.Producer) desc.OcpDocsApiServer {
	return &api{
		repo:    repo,
		prod:    prod,
		flusher: flusher,
	}
}

func (a *api) CreateDocV1(
	ctx context.Context,
	req *desc.CreateDocV1Request,
) (*desc.CreateDocV1Response, error) {
	log.Print("CreateDocV1", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Info().Msgf("Got CreateDocRequest: {name: %s, link: %s, source_link: %s}",
		req.Doc.Name, req.Doc.Link, req.Doc.SourceLink)
	doc := fromMessageNewDoc(req.Doc, 0)
	docId, err := a.repo.AddDoc(ctx, doc)

	if err != nil {
		log.Error().Err(err).Msg("failed to CreateDoc")
		metrics.IncrementCreate(1, "unsuccessful")
		return nil, status.Error(codes.Internal, err.Error())
	}
	a.prod.SendMessage(a.prod.CreateMessage(producer.Created, docId))
	metrics.IncrementCreate(1, "successful")

	log.Info().Msgf("Create doc with id = %d successfully", docId)

	return &desc.CreateDocV1Response{Id: docId}, nil
}

func (a *api) MultiCreateDocsV1(
	ctx context.Context,
	req *desc.MultiCreateDocsV1Request,
) (*desc.MultiCreateDocsV1Response, error) {
	log.Info().Msg("Multi create docs ...")
	span, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateDocsV1")
	defer span.Finish()
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		metrics.IncrementCreate(1, "unsuccessful")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	docs := make([]document.Document, 0, len(req.Docs))
	for _, val := range req.Docs {
		docs = append(docs, fromMessageDoc(val))
	}

	_, idOfCreatedDocs, err := a.flusher.Flush(ctx, docs)

	if err != nil {
		metrics.IncrementCreate(len(docs), "unsuccessful")
		log.Error().Err(err).Msg("failed to multi create docs")
		return nil, status.Error(codes.Internal, err.Error())
	}

	numberOfCreatedDocs := len(idOfCreatedDocs)
	for _, val := range idOfCreatedDocs {
		a.prod.SendMessage(a.prod.CreateMessage(producer.Created, val))
	}
	metrics.IncrementCreate(numberOfCreatedDocs, "successful")
	log.Info().Msgf("MultiCreateDocV1 successful")

	span.SetTag("docs-created", numberOfCreatedDocs)
	return &desc.MultiCreateDocsV1Response{
		DocsAdded: uint64(numberOfCreatedDocs),
	}, nil
}

func (a *api) UpdateDocV1(
	ctx context.Context,
	req *desc.UpdateDocV1Request,
) (*desc.UpdateDocV1Response, error) {
	log.Info().Msgf("Update doc (id: %d) ...", req.Id)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		metrics.IncrementUpdate(1, "unsuccessful")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := a.repo.UpdateDoc(ctx, fromMessageNewDoc(req.Doc, req.Id)); err != nil {
		log.Error().Err(err).Msg("Failed to update doc")
		metrics.IncrementUpdate(1, "unsuccessful")
		return &desc.UpdateDocV1Response{Found: false}, err
	}
	a.prod.SendMessage(a.prod.CreateMessage(producer.Updated, req.Id))
	metrics.IncrementUpdate(1, "successful")
	log.Info().Msgf("Doc was updated")
	return &desc.UpdateDocV1Response{Found: true}, nil
}

func (a *api) ListDocsV1(
	ctx context.Context,
	req *desc.ListDocsV1Request,
) (*desc.ListDocsV1Response, error) {
	log.Print("ListDocsV1", req)
	if err := req.Validate(); err != nil {
		metrics.IncrementRead(1, "unsuccessful")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Info().Msgf("Requesting to list docs: {count: %d, started from: %d}", req.Limit, req.Offset)
	docs, err := a.repo.ListDocs(ctx, req.Limit, req.Offset)
	if err != nil {
		metrics.IncrementRead(1, "unsuccessful")
		log.Error().Err(err).Msg("failed to ListDocs")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	respDocs := make([]*desc.Doc, 0, len(docs))
	for _, doc := range docs {
		respDocs = append(respDocs, toMessage(doc))
		a.prod.SendMessage(a.prod.CreateMessage(producer.Updated, doc.Id))
	}

	metrics.IncrementRead(len(respDocs), "successful")

	return &desc.ListDocsV1Response{Docs: respDocs}, nil
}

func (a *api) DescribeDocV1(
	ctx context.Context,
	req *desc.DescribeDocV1Request,
) (*desc.DescribeDocV1Response, error) {
	log.Print("DescribeDocV1", req)
	log.Info().Msgf("Request describe doc with id = %d", req.Id)
	if err := req.Validate(); err != nil {
		metrics.IncrementRead(1, "unsuccessful")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	doc, err := a.repo.DescribeDoc(ctx, req.Id)

	if err != nil {
		log.Error().Err(err).Msgf("failed to DescribeDoc, id = %d", req.Id)
		metrics.IncrementRead(1, "unsuccessful")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.DescribeDocV1Response{
		Doc: toMessage(*doc),
	}
	a.prod.SendMessage(a.prod.CreateMessage(producer.Described, req.Id))
	metrics.IncrementRead(1, "successful")

	return response, nil
}

func (a *api) RemoveDocV1(
	ctx context.Context,
	req *desc.RemoveDocV1Request,
) (*desc.RemoveDocV1Response, error) {
	log.Print("RemoveDocV1", req)
	log.Info().Msgf("Request delete doc with id = %d", req.Id)
	if err := req.Validate(); err != nil {
		metrics.IncrementDelete(1, "unsuccessful")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.repo.RemoveDoc(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to remove doc, id = %d", req.Id)
		metrics.IncrementDelete(1, "unsuccessful")
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("Doc %d was deleted", req.Id)

	a.prod.SendMessage(a.prod.CreateMessage(producer.Removed, req.Id))
	metrics.IncrementDelete(1, "successful")

	return &desc.RemoveDocV1Response{Found: true}, nil
}
