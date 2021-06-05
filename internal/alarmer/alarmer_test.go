package alarmer_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ocp-docs-api/internal/alarmer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Alarmer", func() {
	var (
		ctrl *gomock.Controller
	)

	AfterEach(func() {
		ctrl.Finish()
	})

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})

	Context("Alarmer loop", func() {
		It("should be closed correctly", func() {
			al := alarmer.New(10 * time.Millisecond)
			al.Init()
			go func() {
				time.Sleep(100 * time.Millisecond)
				al.Close()
			}()
			Eventually(al.Alarm()).Should(BeClosed())
		})
	})
})
