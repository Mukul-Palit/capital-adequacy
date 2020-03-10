package metrics

import (
	"encoding/json"
	"runtime"

	"github.com/gookit/color"
)

//Monitor : Memory Usage monitor
type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}

//NewMonitor : Calculate memory usage
func NewMonitor() {
	var m Monitor
	var rtm runtime.MemStats
	// Read full mem stats
	runtime.ReadMemStats(&rtm)

	// Number of goroutines
	m.NumGoroutine = runtime.NumGoroutine()

	// Misc memory stats
	m.Alloc = bToMb(rtm.Alloc)
	m.TotalAlloc = bToMb(rtm.TotalAlloc)
	m.Sys = bToMb(rtm.Sys)
	m.Mallocs = bToMb(rtm.Mallocs)
	m.Frees = bToMb(rtm.Frees)

	// Live objects = Mallocs - Frees
	m.LiveObjects = m.Mallocs - m.Frees

	// GC Stats
	m.PauseTotalNs = rtm.PauseTotalNs
	m.NumGC = rtm.NumGC

	// Just encode to json and print
	json.MarshalIndent(m, "", " ")
	color.New(color.FgWhite, color.BgLightRed).Printf("Alloc = %v MiB", m.Alloc)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nTotalAlloc = %v MiB", m.TotalAlloc)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nSys = %v MiB", m.Sys)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nMallocs = %v MiB", m.Mallocs)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nFrees = %v MiB", m.Frees)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nLiveObjects = %v", m.Mallocs-m.Frees)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nPauseTotalNs = %v", m.PauseTotalNs)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nNumGC = %v", m.NumGC)
	color.New(color.FgWhite, color.BgLightRed).Printf("\nNumGoroutine = %v\n", m.NumGoroutine)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
