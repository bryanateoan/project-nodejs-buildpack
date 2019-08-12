package hooks

import (
	"runtime"
	"github.com/cloudfoundry/libbuildpack"
	"os"
)

type BryanMemoryStatsHook struct {
	libbuildpack.DefaultHook
	Log *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(BryanMemoryStatsHook{
		Log: logger,
	})
}

func (h BryanMemoryStatsHook) AfterCompile(stager *libbuildpack.Stager) error{
	h.printMemStats()
	return nil
}

/*
*	prints out the current memory statistics during the call.
* 	example from https://golangcode.com/print-the-current-memory-usage/
*	runtime memstats : https://golang.org/src/runtime/mstats.go
*/
func (h BryanMemoryStatsHook) printMemStats() {
	h.Log.Debug("Reading memory statistics")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	h.Log.Debug("Printing memory statistics")
	h.printMemAlloc(&m)
	h.printMemTotalAlloc(&m)
	h.printMemSys(&m)
	h.printMemFrees(&m)
	h.printMemNumGC(&m)
}

func (h BryanMemoryStatsHook) printMemAlloc(m *runtime.MemStats) {
	h.Log.Debug("Printing MemStats Alloc")
	h.Log.Info("Alloc (Allocated heap objects) = %d bytes", m.Alloc)
}

func (h BryanMemoryStatsHook) printMemTotalAlloc(m *runtime.MemStats) {
	h.Log.Debug("Printing MemStats TotalAlloc")
	h.Log.Info("Total Alloc (Cumulative bytes allocated for heap objects) = %d bytes", m.TotalAlloc)
}

func (h BryanMemoryStatsHook) printMemSys(m *runtime.MemStats) {
	h.Log.Debug("Printing MemStats Sys")
	h.Log.Info("Sys (Total bytes of Memory obtained from OS) = %d bytes", m.Sys)
}

func (h BryanMemoryStatsHook) printMemFrees(m *runtime.MemStats) {
	h.Log.Debug("Printing MemStats Frees")
	h.Log.Info("Frees (Cumulative count of heap objects freed) = %d", m.Frees)
}

func (h BryanMemoryStatsHook) printMemNumGC(m *runtime.MemStats) {
	h.Log.Debug("Printing MemStats NumGC")
	h.Log.Info("NumGC (Number of completed GC cyles) = %d", m.NumGC)
}
