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
*	runtime memstats : https://golang.org/src/runtime/mstats.go
*/
func (h BryanSystemStatsHook) printMemStats() {
	h.Log.Debug("Reading memory Statistics")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	h.Log.Info("Alloc (Allocated heap objects) = %d bytes", m.Alloc)
	h.Log.Info("Total Alloc (Cumulative bytes allocated for heap objects) = %d bytes", m.TotalAlloc)
	h.Log.Info("Sys (Total bytes of Memory from OS) = %d bytes", m.Sys)
	h.Log.Info("Frees (Cumulative count of heap objects freed) = %d", m.Frees)
	h.Log.Info("NumGC (Number of completed GC cyles) = %d", m.NumGC)
}

