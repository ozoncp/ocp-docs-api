package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ocp-docs-api/internal/flusher"
	"github.com/ocp-docs-api/internal/mocks"
	"github.com/ocp-docs-api/internal/models/document"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flusher", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		mockRepo *mocks.MockRepo
		docs     []document.Document
		result   []document.Document
		ids      []uint64
		err      error

		f flusher.Flusher

		chunkSize int
	)

	AfterEach(func() {
		ctrl.Finish()
	})

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.New(mockRepo, chunkSize)
		result, ids, err = f.Flush(ctx, docs)
	})

	Context("repo save all tasks", func() {
		BeforeEach(func() {
			chunkSize = 3
			docs = []document.Document{
				{Id: 1, Name: "test1", Link: "www1", SourceLink: "com1"},
				{Id: 2, Name: "test2", Link: "www2", SourceLink: "com2"},
				{Id: 3, Name: "test3", Link: "www3", SourceLink: "com3"},
			}
			mockRepo.EXPECT().AddDocs(gomock.Any(), docs).Return([]uint64{1, 2, 3}, nil)
		})

		It("", func() {
			Expect(len(ids)).Should(Equal(3))
			Expect(result).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("repo save part tasks", func() {
		BeforeEach(func() {
			docs = []document.Document{
				{Id: 1},
				{Id: 2},
				{Id: 3},
				{Id: 4},
			}

			chunkSize = 2
			gomock.InOrder(
				mockRepo.EXPECT().AddDocs(gomock.Any(), []document.Document{{Id: 1}, {Id: 2}}).Return([]uint64{1, 2}, nil),
				mockRepo.EXPECT().AddDocs(gomock.Any(), []document.Document{{Id: 3}, {Id: 4}}).Return([]uint64{}, errors.New("testError")),
			)
		})
		expectedRes := []document.Document{
			{Id: 3},
			{Id: 4},
		}
		It("", func() {
			Expect(result).Should(BeEquivalentTo(expectedRes))
		})
	})

	Context("repo works with empty doc slice", func() {
		BeforeEach(func() {
			chunkSize = 3
			docs = []document.Document{}
		})

		It("", func() {
			Expect(result).Should(BeNil())
		})
	})

	Context("repo works with nil doc slice", func() {
		BeforeEach(func() {
			chunkSize = 3
			docs = nil
		})

		It("", func() {
			Expect(result).Should(BeNil())
		})
	})
})
