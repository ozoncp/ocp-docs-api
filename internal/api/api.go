package api

import (
	"context"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/repo"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOcpDocsApiServer
	repo repo.Repo
}

func toMessage(doc document.Document) *desc.Doc {
	return &desc.Doc{
		Id: doc.Id,
		Name: doc.Name,
		Link: doc.Link,
		SourceLink: doc.SourceLink,
	}
}

func NewDocsApi(repo repo.Repo) desc.OcpDocsApiServer {
	return &api{repo: repo}
}

func (a *api) ListDocsV1(
	ctx context.Context,
	req *desc.ListDocsV1Request,
) (*desc.ListDocsV1Response, error) {
	log.Print("ListDocsV1", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Info().Msgf("Requesting to list docs: {count: %d, started from: %d}", req.Limit, req.Offset)
	docs, err := a.repo.ListDocs(ctx, req.Limit, req.Offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to ListDocs")
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var respDocs []*desc.Doc

	for _, doc := range docs {
		respDocs = append(respDocs, toMessage(doc))
	}
	return &desc.ListDocsV1Response{Docs: respDocs}, nil
}

func (a *api) DescribeDocV1(
	ctx context.Context,
	req *desc.DescribeDocV1Request,
) (*desc.DescribeDocV1Response, error) {
	log.Print("DescribeDocV1", req)
	log.Info().Msgf("Request describe doc with id = %d", req.Id)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	doc, err := a.repo.DescribeDoc(ctx, req.Id)

	if err != nil {
		log.Error().Err(err).Msgf("failed to DescribeDoc, id = %d", req.Id)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &desc.DescribeDocV1Response{
		Doc: &desc.Doc{
			Id: doc.Id,
			Name: doc.Name,
			Link: doc.Link,
			SourceLink: doc.SourceLink,
		},
	}
	return response, nil
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
																					req.Name, req.Link, req.SourceLink)
	doc := document.Document{
					Name: req.Name,
					Link: req.Link,
					SourceLink: req.SourceLink,
	}
	docId, err := a.repo.AddDoc(ctx, doc)

	if err != nil {
		log.Error().Err(err).Msg("failed to CreateDoc")
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("Create doc with id = %d successfully", docId)

	return &desc.CreateDocV1Response{Id:docId}, nil
}

func (a *api) RemoveDocV1(
	ctx context.Context,
	req *desc.RemoveDocV1Request,
) (*desc.RemoveDocV1Response, error) {
	log.Print("RemoveDocV1", req)
	log.Info().Msgf("Request delete doc with id = %d", req.Id)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.repo.RemoveDoc(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msgf("failed to remove doc, id = %d", req.Id)
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("Doc %d was deleted", req.Id)
	return &desc.RemoveDocV1Response{Found: true}, nil
}