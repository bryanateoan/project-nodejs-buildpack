package hooks_test

import (
	"bytes"
	"github.com/cloudfoundry/libbuildpack"
	"github.com/cloudfoundry/nodejs-buildpack/src/nodejs/hooks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("bryanSystemStatsHook", func() {
	var(
		err error
		logger *libbuildpack.Logger
		stager *libbuildpack.Stager
		buffer *bytes.Buffer
		bryan hooks.BryanSystemStatsHook
	)

	BeforeEach(func() {
		buffer = new(bytes.Buffer)
		logger = libbuildpack.NewLogger(buffer)
		bryan = hooks.BryanSystemStatsHook {
			Log: logger,
		}
	})

	JustBeforeEach(func() {
		args := []string{"", "", "", ""}
		stager = libbuildpack.NewStager(args, logger, &libbuildpack.Manifest{})
	})

	Describe("AfterCompile", func() {
		Context("Memory Allocation Statistics are printed", func() {
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

		})
	})

})