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
			It("prints out Alloc with the allocated memory",  func() {
				err = bryan.AfterCompile(stager)
				Expect(err).To(BeNil())
				Expect(buffer.String()).To(ContainSubstring("Alloc ="))

			})
		})
	})

})