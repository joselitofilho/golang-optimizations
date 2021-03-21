package memoryleak

import "testing"

func BenchmarkKindofLeakingCausedByInteriorPointer(b *testing.B) {
	CausedByInteriorPointer()
	// keep using p in the rest of the app,
	// never using data from the struct

	// runtime.GC()
	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Printf("  end %d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", m.Mallocs, m.Frees, m.Alloc, m.TotalAlloc, m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, m.NumGC, m.NumForcedGC)
}
