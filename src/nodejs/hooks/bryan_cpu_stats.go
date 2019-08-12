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
	cpuCountLogical,cpuErrorLogical := cpu.Counts(true)
	if(cpuErrorLogical != nil) {
		return cpuErrorLogical
	}

	cpuCountPhysical,cpuErrorPhysical := cpu.Counts(false)
	if(cpuErrorPhysical != nil) {
		return cpuErrorPhysical
	}

	h.Log.Info("Logical CPU count: %d", cpuCountLogical)
	h.Log.Info("Physical CPU count: %d", cpuCountPhysical)
	return nil
}