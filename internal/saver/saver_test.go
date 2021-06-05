package saver_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ocp-docs-api/internal/alarmer"
	"github.com/ocp-docs-api/internal/mocks"
	"github.com/ocp-docs-api/internal/models/document"
	"github.com/ocp-docs-api/internal/saver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctrl *gomock.Controller
		mockFlusher *mocks.MockFlusher
	)

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("New Call", func(){
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFlusher = mocks.NewMockFlusher(ctrl)
		})

		Context("Constructor with valid args", func() {
			It("should return valid object", func() {
				a := alarmer.New(time.Second)
				got := saver.New(1, mockFlusher, a, saver.DropAll)
				Ω(got).ShouldNot(BeNil())
			})
		})
	})

	Describe("Functionality", func(){
		var (
			alarmer     *alarmerStub
		)
		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
			mockFlusher = mocks.NewMockFlusher(ctrl)
			alarmer = &alarmerStub{make(chan struct{})}
		})

		Context("Save nothing before closing ", func() {
			It("should flush everything on alarm and close shouldn't flush anything", func() {
				docs := []document.Document{
					{Id: 1},
					{Id: 2},
				}
				var wg sync.WaitGroup
				wg.Add(1)
				defer wg.Wait()
				saver := saver.New(len(docs) + 1, mockFlusher, alarmer, saver.DropAll)

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(gomock.Eq(docs)),
					mockFlusher.EXPECT().Flush(gomock.Eq(docs[:0])).Do(func(entities []document.Document) { wg.Done()}),
				)
				saver.Init()

				for i := 0; i < len(docs); i++ {
					saver.Save(docs[i])
				}
				alarmer.alarm()
				saver.Close()
			})
		})

		Context("Save smth before closing ", func() {
			It("", func() {
				docs := []document.Document{
					{Id: 1},
					{Id: 2},
					{Id: 3},
					{Id: 4},
				}
				var wg sync.WaitGroup
				wg.Add(1)
				defer wg.Wait()
				saver := saver.New(len(docs) + 1, mockFlusher, alarmer, saver.DropAll)

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(gomock.Eq(docs[:2])),
					mockFlusher.EXPECT().Flush(gomock.Eq(docs[2:])).Do(func(entities []document.Document) { wg.Done()}),
				)
				saver.Init()

				for i := 0; i < 2; i++ {
					saver.Save(docs[i])
				}
				alarmer.alarm()
				for i := 2; i < len(docs); i++ {
					saver.Save(docs[i])
				}
				saver.Close()
			})
		})

		Context("Save more than capacity", func() {
			It("should drop everything", func() {
				docs := []document.Document{
					{Id: 1},
					{Id: 2},
					{Id: 3},
					{Id: 4},
				}

				newDocs := []document.Document{
					{Id: 5},
					{Id: 6},
				}

				var wg sync.WaitGroup
				wg.Add(1)
				defer wg.Wait()

				saver := saver.New(len(docs), mockFlusher, alarmer, saver.DropAll)

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(gomock.Eq(newDocs)),
					mockFlusher.EXPECT().Flush(gomock.Eq(newDocs[:0])).Do(func(entities []document.Document) { wg.Done() }),
				)

				saver.Init()

				for i := 0; i < len(docs); i++ {
					saver.Save(docs[i])
				}

				for i := 0; i < len(newDocs); i++ {
					saver.Save(newDocs[i])
				}

				alarmer.alarm()
				saver.Close()
			})

			It("should drop first", func() {
				docs := []document.Document{
					{Id: 1},
					{Id: 2},
					{Id: 3},
				}

				newDocs := []document.Document{
					{Id: 2},
					{Id: 3},
					{Id: 4},
				}

				var wg sync.WaitGroup
				wg.Add(1)
				defer wg.Wait()

				saver := saver.New(len(docs), mockFlusher, alarmer, saver.DropOne)

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(gomock.Eq(newDocs)),
					mockFlusher.EXPECT().Flush(gomock.Eq(newDocs[:0])).Do(func(entities []document.Document) { wg.Done() }),
				)

				saver.Init()

				for i := 0; i < len(docs); i++ {
					saver.Save(docs[i])
				}
				saver.Save(document.Document{Id: 4})
				alarmer.alarm()
				saver.Close()
			})
		})

		Context("flush has failed ", func() {
			It("", func() {
				docs := []document.Document{
					{Id: 1},
					{Id: 2},
					{Id: 3},
					{Id: 4},
				}
				var wg sync.WaitGroup
				wg.Add(1)
				defer wg.Wait()
				saver := saver.New(len(docs) + 1, mockFlusher, alarmer, saver.DropAll)

				gomock.InOrder(
					mockFlusher.EXPECT().Flush(gomock.Eq(docs)).Return([]document.Document{{Id: 3}, {Id: 4}}),
					mockFlusher.EXPECT().Flush(gomock.Eq([]document.Document{{Id: 3}, {Id: 4}})).Do(func(entities []document.Document) { wg.Done()}),
				)
				saver.Init()

				for i := 0; i < len(docs); i++ {
					saver.Save(docs[i])
				}
				alarmer.alarm()
				saver.Close()
			})
		})
	})

})

type alarmerStub struct {
	alarms chan struct{}
}

func (a *alarmerStub) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarmerStub) alarm() {
	a.alarms <- struct{}{}
}

func (a * alarmerStub) Close() {
}

func (a * alarmerStub) Init() {
}