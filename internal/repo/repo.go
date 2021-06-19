package repo

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

type Repo interface {
	AddDoc(ctx context.Context, doc document.Document) (uint64, error)
	AddDocs(ctx context.Context, docs []document.Document) (uint64, error)
	RemoveDoc(ctx context.Context, docId uint64) error
	DescribeDoc(ctx context.Context, docId uint64) (*document.Document, error)
	ListDocs(ctx context.Context, limit, offset uint64) ([]document.Document, error)
	UpdateDoc(ctx context.Context, doc document.Document) error
}

const (
	tableName = "docs"
)

type repo struct {
	db sqlx.DB
	chunkSize int
}

func New(db sqlx.DB, chunkSize int) Repo {
	return &repo{
		db: db,
		chunkSize: chunkSize,
	}
}

func (r *repo) AddDoc(ctx context.Context, doc document.Document) (uint64, error) {
	query := sq.Insert(tableName).
		Columns("name", "link", "source_link").
		Values(doc.Name, doc.Link, doc.SourceLink).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&doc.Id)

	if err != nil {
		return 0, err
	}

	return doc.Id, nil
}

func (r *repo) RemoveDoc(ctx context.Context, docId uint64) error {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": docId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)
	result, err := query.ExecContext(ctx)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("doc not found")
	}

	return nil
}

func (r *repo) DescribeDoc(ctx context.Context, docId uint64) (*document.Document, error) {
	query := sq.Select("id", "name", "link", "source_link").
		From(tableName).
		Where(sq.Eq{"id": docId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	doc := &document.Document{}
	return doc, query.QueryRowContext(ctx).Scan(&doc.Id, &doc.Name, &doc.Link, &doc.SourceLink)
}

func (r *repo) ListDocs(ctx context.Context, limit, offset uint64) ([]document.Document, error) {
	query := sq.Select("id", "name", "link", "source_link").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)
	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	listDocs := make([]document.Document, 0)
	defer rows.Close()
	for rows.Next() {
		var doc document.Document
		err := rows.Scan(&doc.Id, &doc.Name, &doc.Link, &doc.SourceLink)
		if err != nil {
			continue
		}
		listDocs = append(listDocs, doc)
	}
	return listDocs, nil
}

func (r *repo) AddDocs(ctx context.Context, docs []document.Document) (uint64, error) {
	chunks, err := utils.SplitDocumentSlice(docs, r.chunkSize)
	var numberOfDocsCreated int64 = 0

	if err != nil {
	 	return uint64(numberOfDocsCreated), err
	}

	for i, val := range chunks {
		err := func() error {
			spanName := "Add_docs" + fmt.Sprintf("%v", i)
			span, _ := opentracing.StartSpanFromContext(ctx, spanName)
			defer span.Finish()

			query := sq.Insert(tableName).
				Columns("name", "link", "source_link").
				RunWith(r.db).
				PlaceholderFormat(sq.Dollar)

			for _, doc := range val {
				query = query.Values(doc.Name, doc.Link, doc.SourceLink)
			}

			result, err := query.ExecContext(ctx)

			if err != nil {
				return err
			}

			rowsAffected, err := result.RowsAffected()

			if err != nil {
				return err
			}

			numberOfDocsCreated += rowsAffected
			span.SetTag("docs-added-in-db", numberOfDocsCreated)
			return nil
		}()

		if err != nil {
			return uint64(numberOfDocsCreated), err
		}
	}

	return uint64(numberOfDocsCreated), nil
}

func (r *repo) UpdateDoc(ctx context.Context, doc document.Document) error {
	query := sq.Update(tableName).
			Where(sq.Eq{"id": doc.Id}).
		    Set("name", doc.Name).
			Set("link", doc.Link).
			Set("source_link", doc.SourceLink).
		    RunWith(r.db).
		    PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("doc not found")
	}

	return nil
}
