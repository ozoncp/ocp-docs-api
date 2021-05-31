package flusher_test

import (
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
		ctrl *gomock.Controller

		mockRepo *mocks.MockRepo
		docs []document.Document
		result []document.Document

		f flusher.Flusher

		chunkSize int
	)

	AfterEach(func() {
		ctrl.Finish()
	})

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func(){
		f = flusher.New(mockRepo, chunkSize)
		result = f.Flush(docs)
	})

	Context("repo save all tasks", func() {
		BeforeEach(func() {
			chunkSize = 3
			//docs = []document.Document{}
			docs = []document.Document{
				{Id: 1},
				{Id: 2},
				{Id: 3},
			}
			mockRepo.EXPECT().AddDocs(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(result).Should(BeNil())
		})
	})

	Context("repo works with empty doc slice", func() {
		BeforeEach(func() {
			chunkSize = 3
			docs = []document.Document{}
			mockRepo.EXPECT().AddDocs(gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(result).Should(BeNil())
		})
	})

	Context("repo save part tasks", func() {
		var (
			halfSize int
		)
		BeforeEach(func() {
			docs = []document.Document{
				{Id: 1},
				{Id: 2},
				{Id: 3},
				{Id: 4},
			}
			halfSize = int(len(docs) / 2)
			chunkSize = halfSize

			gomock.InOrder(
				mockRepo.EXPECT().AddDocs(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockRepo.EXPECT().AddDocs(gomock.Len(len(docs)-chunkSize)).Return(errors.New("testError")),
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
})
