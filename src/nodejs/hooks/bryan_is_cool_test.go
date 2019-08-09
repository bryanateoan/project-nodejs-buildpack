package hooks_test

import (
	"bytes"
	"github.com/cloudfoundry/libbuildpack"
	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bryanHook", func() {
	var (
		err error
		logger *libbuildpack.Logger
		stager *libbuildpack.Stager
		buffer *bytes.Buffer
		bryan hooks.BryanHook
	)
	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		bryan = hooks.BryanHook{
			Message:  "Bryan is Cool",
			Log: logger,
		}
	})
	
	JustBeforeEach(func() {
		args := []string{"", "", "", ""}
		stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
	})

	Describe("AfterCompile", func() {
		Context("Message is said", func() {
			It("says the message", func() {
				err = bryan.AfterCompile(stager)
				Expect(err).To(BeNil())
				Expect(buffer.String()).To(ContainSubstring("Bryan is Cool"))
			})
		})
	})
})
