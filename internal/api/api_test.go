package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/ocp-docs-api/internal/api"
	"github.com/ocp-docs-api/internal/mocks"
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
		dataProducerMock *mocks.MockProducer
	)
	BeforeEach(func() {
		ctx = context.Background()
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		ctrl := gomock.NewController(GinkgoT())
		dataProducerMock = mocks.NewMockProducer(ctrl)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("test api functions", func() {
		BeforeEach(func() {
			//prod, err = producer.NewProducer("TestOcpDocsApiCreate")
			Expect(err).Should(BeNil())
			testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
		})

		Context("test create functions", func(){
			It("Test create doc", func() {
				request := &desc.CreateDocV1Request{
					Doc: &desc.NewDoc{
						Name:       "testName",
						Link:       "www",
						SourceLink: "com",
					},
				}

				mock.ExpectQuery("INSERT INTO docs").
					WithArgs(request.Doc.Name, request.Doc.Link, request.Doc.SourceLink).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				Expect(testApi).ShouldNot(BeNil())
				dataProducerMock.EXPECT().SendMessage("CreateDocV1 succesful")
				response, err := testApi.CreateDocV1(ctx, request)
				Expect(err).Should(BeNil())
				Expect(response.Id).Should(BeEquivalentTo(1))
			})

			It("Test incorrect create docs", func() {
				request := &desc.CreateDocV1Request{
					Doc: &desc.NewDoc{
						Name:       "testName",
						Link:       "www",
						SourceLink: "com",
					},
				}

				mock.ExpectQuery("INSERT INTO docs").
					WithArgs(request.Doc.Name, request.Doc.Link, request.Doc.SourceLink).
					WillReturnError(errors.New("failed to execute sql request"))
				Expect(testApi).ShouldNot(BeNil())

				response, err := testApi.CreateDocV1(ctx, request)
				Expect(err).ShouldNot(BeNil())
				Expect(response).Should(BeNil())
			})
		})

		Context("test multi-create", func(){
			BeforeEach(func() {
				testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
			})

			It("Test correct multi-creation", func(){
				request := &desc.MultiCreateDocsV1Request{
					Docs: []*desc.Doc{
						{Id: 1, Name: "test1", Link: "www1", SourceLink: "com1"},
						{Id: 2, Name: "test2", Link: "www2", SourceLink: "com2"},
						{Id: 3, Name: "test3", Link: "www3", SourceLink: "com3"}},
				}
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1).AddRow(2).AddRow(3)
				mock.ExpectQuery("INSERT INTO docs").
					WithArgs("test1", "www1", "com1",
						"test2", "www2", "com2",
						"test3", "www3", "com3").WillReturnRows(rows)
				dataProducerMock.EXPECT().SendMessage("MultiCreateDocV1 successful")

				response, err := testApi.MultiCreateDocsV1(ctx, request)
				Expect(err).Should(BeNil())
				Expect(response.DocsAdded).Should(Equal(uint64(3)))
			})

			It("Test incorrect multi-creation", func(){
				request := &desc.MultiCreateDocsV1Request{
					Docs: []*desc.Doc{
						{Id: 1, Name: "test1", Link: "www1", SourceLink: "com1"},
						{Id: 2, Name: "test2", Link: "www2", SourceLink: "com2"},
						{Id: 3, Name: "test3", Link: "www3", SourceLink: "com3"}},
				}

				mock.ExpectQuery("INSERT INTO docs").
					WithArgs("test1", "www1", "com1",
						"test2", "www2", "com2",
						"test3", "www3", "com3").WillReturnError(errors.New("failed to execute sql request"))

				response, err := testApi.MultiCreateDocsV1(ctx, request)
				Expect(err).ShouldNot(BeNil())
				Expect(response).Should(BeNil())
			})
		})

		Context("Update docs", func(){
			BeforeEach(func() {
				testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
			})

			It("Test update doc", func(){
				doc := &desc.NewDoc{
					Name: "test1",
					Link: "www1",
					SourceLink: "com1",
				}

				request := &desc.UpdateDocV1Request{
					Id: 1,
					Doc : doc,
				}

				mock.ExpectExec("UPDATE docs").
					WithArgs("test1", "www1", "com1", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				dataProducerMock.EXPECT().SendMessage("UpdateDocV1 successful")

				response, err := testApi.UpdateDocV1(ctx, request)
				Expect(err).Should(BeNil())
				Expect(response.Found).Should(Equal(true))
			})

			It("Test incorrect update doc", func(){
				doc := &desc.NewDoc{
					Name: "test1",
					Link: "www1",
					SourceLink: "com1",
				}

				request := &desc.UpdateDocV1Request{
					Id: 1,
					Doc : doc,
				}

				mock.ExpectExec("UPDATE docs").
					WithArgs("test1", "www1", "com1", 1).
					WillReturnError(errors.New("failed to execute sql request"))

				response, err := testApi.UpdateDocV1(ctx, request)
				Expect(err).ShouldNot(BeNil())
				Expect(response.Found).Should(Equal(false))
			})
		})

		Context("Test Remove doc", func(){
			BeforeEach(func() {
				//prod, err = producer.NewProducer("TestOcpDocsApiRemove")
				//Expect(err).Should(BeNil())
				testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
			})

			It("Test remove doc", func() {
				request := &desc.RemoveDocV1Request{
					Id: 1,
				}

				mock.ExpectExec("DELETE FROM docs").
					WithArgs(request.Id).
					WillReturnResult(sqlmock.NewResult(0, 1))
				dataProducerMock.EXPECT().SendMessage("RemoveDocV1 successful")

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
		})

		Context("Describe doc", func(){
			BeforeEach(func() {
				//prod, err = producer.NewProducer("TestOcpDocsApiDescribe")
				//Expect(err).Should(BeNil())
				testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
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
				dataProducerMock.EXPECT().SendMessage("DescribeDocV1 successful")

				response, err := testApi.DescribeDocV1(ctx, request)
				Expect(err).Should(BeNil())
				Expect(response.Doc).Should(Equal(&desc.Doc{Id: 1, Name: "testName", Link: "www", SourceLink: "com"}))
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
		})

		Context("List doc", func(){
			BeforeEach(func() {
				//prod, err = producer.NewProducer("TestOcpDocsApiList")
				//Expect(err).Should(BeNil())
				testApi = api.NewDocsApi(repo.New(*sqlxDB), dataProducerMock)
			})

			It("Test list doc", func() {
				request := &desc.ListDocsV1Request{
					Limit:  3,
					Offset: 0,
				}
				docs := []document.Document{
					{Id: 1, Name: "test1", Link: "link1", SourceLink: "srcLink1"},
					{Id: 2, Name: "test2", Link: "link2", SourceLink: "srcLink2"},
				}
				query := fmt.Sprintf("SELECT (.+) FROM docs LIMIT %d OFFSET %d", request.Limit, request.Offset)
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.
						NewRows([]string{"id", "name", "link", "source_link"}).
						AddRow(docs[0].Id, docs[0].Name, docs[0].Link, docs[0].SourceLink).
						AddRow(docs[1].Id, docs[1].Name, docs[1].Link, docs[1].SourceLink))
				dataProducerMock.EXPECT().SendMessage("ListDocsV1 successful")

				response, err := testApi.ListDocsV1(ctx, request)

				Expect(err).Should(BeNil())
				Expect(response).Should(Equal(
					&desc.ListDocsV1Response {
						Docs: []*desc.Doc{
							{Id: 1, Name: "test1", Link: "link1", SourceLink: "srcLink1"},
							{Id: 2, Name: "test2", Link: "link2", SourceLink: "srcLink2"},
						},
					},
				))
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
})
