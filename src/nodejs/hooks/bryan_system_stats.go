package hooks

import (
	"runtime"
	"github.com/cloudfoundry/libbuildpack"
	"os"
)

type BryanSystemStatsHook struct {
	libbuildpack.DefaultHook
	Log *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(BryanSystemStatsHook{
		Log: logger,
	})
}

func (h BryanSystemStatsHook) AfterCompile(stager *libbuildpack.Stager) error{
	h.printMemStats()
	return nil
}

/*
*	prints out the current memory allocator statistics during the call.
* 	example from https://golangcode.com/print-the-current-memory-usage/
*/
func (h BryanSystemStatsHook) printMemStats() {
	h.Log.Debug("Reading memory Statistics")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	h.Log.Info("Alloc = %v b", m.Alloc)

} 