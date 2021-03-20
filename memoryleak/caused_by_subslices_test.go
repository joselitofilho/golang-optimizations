package memoryleak

import (
	"testing"
)

func BenchmarkKindofLeakingCausedBySubslices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CausedBySubslices()
	}

	// runtime.GC()
	// var m runtime.MemStats
	// runtime.ReadMemStats(&m)
	// fmt.Printf("  end %d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", m.Mallocs, m.Frees, m.Alloc, m.TotalAlloc, m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, m.NumGC, m.NumForcedGC)
}
