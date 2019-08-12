package hooks_test

import (
	"os"
	"bytes"
	"github.com/cloudfoundry/libbuildpack"
	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bryanCpuStatsHook", func() {
	var(
		err error
		logger *libbuildpack.Logger
		stager *libbuildpack.Stager
		buffer *bytes.Buffer
		bryan hooks.BryanCpuStatsHook
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		bryan = hooks.BryanCpuStatsHook {
			Log: logger,
		}
	})

	JustBeforeEach(func() {
		args := []string{"", "", "", ""}
		stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
	})

	Describe("AfterCompile", func() {
		var(
			oldBpDebug string
		)
		Context("Cpu statistics are printed", func() {
			BeforeEach(func() {
				err = bryan.AfterCompile(stager)
			})

			It("Did not error", func() {
				Expect(err).To(BeNil())
			})
			It("Printed the Logical cores", func() {
				Expect(buffer.String()).To(ContainSubstring("Logical CPU count: "))
			})
			It("Printed the Physical cores", func() {
				Expect(buffer.String()).To(ContainSubstring("Physical CPU count: "))
			})
			It("Printed the Header", func() {
				Expect(buffer.String()).To(ContainSubstring("===CPU Statistics==="))
			})
			It("Printed the Footer", func() {
				Expect(buffer.String()).To(ContainSubstring("=========="))
			})

		})

		Context("Debug mode is true", func() {
			BeforeEach(func() {
				oldBpDebug = os.Getenv("BP_DEBUG")
				os.Setenv("BP_DEBUG", "TRUE")
				err = bryan.AfterCompile(stager)
			})

			AfterEach(func() {
				os.Setenv("BP_DEBUG", oldBpDebug)
			})

			It("Did not error", func() {
				Expect(err).To(BeNil())
			})
			It("Printed debug message for logical cores", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing Logical cpu count"))
			})
			It("Printed debug message for physical cores", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing Physical cpu count"))
			})
			It("Printed debug message for header", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing header"))
			})
			It("Printed debug message for footer", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing footer"))
			})
		})
	})
})