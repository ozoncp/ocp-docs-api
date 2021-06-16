package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ocp-docs-api/internal/api"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/repo"
	desc "github.com/ocp-docs-api/pkg/ocp-docs-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	var (
		ctx     context.Context
		testApi desc.OcpDocsApiServer
		mock    sqlmock.Sqlmock
		db      *sql.DB
		sqlxDB  *sqlx.DB
		err     error
	)
	BeforeEach(func() {
		ctx = context.Background()
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")
		testApi = api.NewDocsApi(repo.New(*sqlxDB))
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("test api functions", func() {

		It("Test create doc", func() {
			request := &desc.CreateDocV1Request{
				Name:       "testName",
				Link:       "www",
				SourceLink: "com",
			}
			mock.ExpectQuery("INSERT INTO docs").
				WithArgs(request.Name, request.Link, request.SourceLink).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.CreateDocV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Id).Should(BeEquivalentTo(1))
		})

		It("Test incorrect create docs", func() {
			request := &desc.CreateDocV1Request{
				Name:       "testName",
				Link:       "www",
				SourceLink: "com",
			}

			mock.ExpectQuery("INSERT INTO docs").
				WithArgs(request.Name, request.Link, request.SourceLink).
				WillReturnError(errors.New("failed to execute sql request"))

			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.CreateDocV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
			Expect(response).Should(BeNil())
		})

		It("Test remove doc", func() {
			request := &desc.RemoveDocV1Request{
				Id: 1,
			}

			mock.ExpectExec("DELETE FROM docs").
				WithArgs(request.Id).
				WillReturnResult(sqlmock.NewResult(0, 1))

			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.RemoveDocV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Found).Should(BeEquivalentTo(true))
		})

		It("Test incorrect remove doc", func() {
			request := &desc.RemoveDocV1Request{
				Id: 1,
			}

			mock.ExpectExec("DELETE FROM docs").
				WithArgs(request.Id).
				WillReturnError(errors.New("failed to remove doc"))

			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.RemoveDocV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
			Expect(response).Should(BeNil())
		})

		It("Test describe docs", func() {
			request := &desc.DescribeDocV1Request{
				Id: 1,
			}

			mock.ExpectQuery("SELECT (.+) FROM docs WHERE").
				WithArgs(request.Id).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "name", "link", "source_link"}).
					AddRow(1, "testName", "www", "com"))

			response, err := testApi.DescribeDocV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Doc.Id).Should(Equal(uint64(1)))
			Expect(response.Doc.Name).Should(Equal("testName"))
			Expect(response.Doc.Link).Should(Equal("www"))
			Expect(response.Doc.SourceLink).Should(Equal("com"))
		})

		It("Test incorrect describe doc", func() {
			request := &desc.DescribeDocV1Request{
				Id: 1,
			}

			mock.ExpectQuery("SELECT (.+) FROM docs WHERE").
				WithArgs(request.Id).
				WillReturnError(errors.New("can't describe doc"))

			response, err := testApi.DescribeDocV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
			Expect(response).Should(BeNil())
		})

		It("Test list doc", func() {
			request := &desc.ListDocsV1Request{
				Limit:  3,
				Offset: 0,
			}
			docs := []document.Document{
				{Id: 1, Name: "test1", Link: "link1", SourceLink: "srcLink1"},
				{Id: 2, Name: "test2", Link: "link3", SourceLink: "srcLink3"},
			}
			query := fmt.Sprintf("SELECT (.+) FROM docs LIMIT %d OFFSET %d", request.Limit, request.Offset)
			mock.ExpectQuery(query).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "name", "link", "source_link"}).
					AddRow(docs[0].Id, docs[0].Name, docs[0].Link, docs[0].SourceLink).
					AddRow(docs[1].Id, docs[1].Name, docs[1].Link, docs[1].SourceLink))

			response, err := testApi.ListDocsV1(ctx, request)

			Expect(err).Should(BeNil())
			Expect(response.Docs[0].Id).Should(Equal(docs[0].Id))
			Expect(response.Docs[0].Name).Should(Equal(docs[0].Name))
			Expect(response.Docs[0].Link).Should(Equal(docs[0].Link))
			Expect(response.Docs[0].SourceLink).Should(Equal(docs[0].SourceLink))

			Expect(response.Docs[1].Id).Should(Equal(docs[1].Id))
			Expect(response.Docs[1].Name).Should(Equal(docs[1].Name))
			Expect(response.Docs[1].Link).Should(Equal(docs[1].Link))
			Expect(response.Docs[1].SourceLink).Should(Equal(docs[1].SourceLink))
		})

		It("Test incorrect list doc", func() {
			request := &desc.ListDocsV1Request{
				Limit:  3,
				Offset: 0,
			}

			query := fmt.Sprintf("SELECT (.+) FROM docs LIMIT %d OFFSET %d", request.Limit, request.Offset)
			mock.ExpectQuery(query).
				WillReturnError(errors.New("can't provide list of docs"))

			response, err := testApi.ListDocsV1(ctx, request)

			Expect(err).ShouldNot(BeNil())
			Expect(response).Should(BeNil())
		})
	})
})
