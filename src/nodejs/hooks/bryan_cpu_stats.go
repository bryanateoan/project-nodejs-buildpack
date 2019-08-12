package hooks

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/cloudfoundry/libbuildpack"
	"os"
)

type BryanCpuStatsHook struct {
	libbuildpack.DefaultHook
	Log *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(BryanCpuStatsHook{
		Log: logger,
	})
}

func (h BryanCpuStatsHook) AfterCompile(stager *libbuildpack.Stager) error{
	error := h.printCpuStats()
	return error
}

func (h BryanCpuStatsHook) printCpuStats() error {
	headerError := h.printHeader()
	if(headerError != nil) {
		return headerError
	}

	logicalError := h.printLogicalCpuCount()
	if(logicalError != nil) {
		return logicalError
	}

	physicalError := h.printPhysicalCpuCount()
	if(physicalError != nil) {
		return physicalError
	}

	footerError := h.printFooter()
	if(footerError != nil) {
		return footerError
	}

	return nil

}

func (h BryanCpuStatsHook) printLogicalCpuCount() error {
	h.Log.Debug("Printing Logical cpu count")
	cpuCountLogical,cpuErrorLogical := cpu.Counts(true)
	if(cpuErrorLogical != nil) {
		return cpuErrorLogical
	}

	h.Log.Info("Logical CPU count: %d", cpuCountLogical)
	return nil
}

func (h BryanCpuStatsHook) printPhysicalCpuCount() error {
	h.Log.Debug("Printing Physical cpu count")
	cpuCountPhysical,cpuErrorPhysical := cpu.Counts(false)
	if(cpuErrorPhysical != nil) {
		return cpuErrorPhysical
	}

	h.Log.Info("Physical CPU count: %d", cpuCountPhysical)
	return nil
}

func (h BryanCpuStatsHook) printHeader() error {
	h.Log.Debug("Printing header")
	h.Log.Info("===CPU Statistics===")
	return nil
}

func (h BryanCpuStatsHook) printFooter() error {
	h.Log.Debug("Printing footer")
	h.Log.Info("==========")
	return nil
}