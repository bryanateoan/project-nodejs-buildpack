package hooks_test

import (
	"os"
	"bytes"
	"github.com/cloudfoundry/libbuildpack"
	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bryanMemoryStatsHook", func() {
	var(
		err error
		logger *libbuildpack.Logger
		stager *libbuildpack.Stager
		buffer *bytes.Buffer
		bryan hooks.BryanMemoryStatsHook
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		bryan = hooks.BryanMemoryStatsHook {
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
		Context("Memory statistics are printed", func() {
			BeforeEach(func() {
				err = bryan.AfterCompile(stager)
			})

			It("did not return any error", func() {
				Expect(err).To(BeNil())
			})
			It("prints out Alloc with the allocated memory",  func() {
				Expect(buffer.String()).To(ContainSubstring("Alloc (Allocated heap objects) ="))

			})
			It("prints out Total Alloc", func() {
				Expect(buffer.String()).To(ContainSubstring("Total Alloc (Cumulative bytes allocated for heap objects) ="))
			})
			It("prints out Sys", func() {
				Expect(buffer.String()).To(ContainSubstring("Sys (Total bytes of Memory obtained from OS) ="))
			})
			It("prints out Frees", func() {
				Expect(buffer.String()).To(ContainSubstring("Frees (Cumulative count of heap objects freed) ="))
			})
			It("prints out NumGC", func() {
				Expect(buffer.String()).To(ContainSubstring("NumGC (Number of completed GC cyles) ="))
			})
			It("prints header", func() {
				Expect(buffer.String()).To(ContainSubstring("===Memory Statistics==="))
			})
			It("prints footer", func() {
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
			It("Printed out the debug message for Alloc", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing MemStats Alloc"))
			})
			It("Printed out the debug message for Total Alloc", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing MemStats TotalAlloc"))
			})
			It("Printed out the debug message for Sys", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing MemStats Sys"))
			})
			It("Printed out the debug message for Frees", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing MemStats Frees"))
			})
			It("Printed out the debug message for NumGC", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing MemStats NumGC"))
			})
			It("Printed out the debug message for header", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing Header"))
			})
			It("Printed out the debug message for footer", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing Footer"))
			})
			It("Printed out the debug message for reading memory", func() {
				Expect(buffer.String()).To(ContainSubstring("Reading memory statistics"))
			})
			It("Printed out the debug message for printing memory", func() {
				Expect(buffer.String()).To(ContainSubstring("Printing memory statistics"))
			})
		})
	})

})