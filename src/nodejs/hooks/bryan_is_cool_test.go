package hooks_test

import (
	"os"
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
		Context("Message is valid", func() {
			It("says the message", func() {
				err = bryan.AfterCompile(stager)
				Expect(err).To(BeNil())
				Expect(buffer.String()).To(ContainSubstring("Bryan is Cool"))
			})
		})

		Context("Message is empty", func() {
			BeforeEach(func() {
				bryan = hooks.BryanHook{
					Message: "",
					Log: logger,
				}
			})
			It("returns an error ", func() {
				err = bryan.AfterCompile(stager)
				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError("no message to print"))
				Expect(buffer.String()).To(ContainSubstring("Failing build..."))
				Expect(buffer.String()).To(ContainSubstring("ERROR"))
			})
		})

		Context("Demo the message, logger is stdout", func() {
			BeforeEach(func() {
				logger = libbuildpack.NewLogger(os.Stdout)
				bryan = hooks.BryanHook{
					Message: "Bryan is Cool",
					Log: logger,
				}
			})
			JustBeforeEach(func() {
				args := []string{"", "", "", ""}
				stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
			})
			It("Demo message", func() {
				err = bryan.AfterCompile(stager)
				Expect(err).To(BeNil())
			})
		})
	})
})
