package memoryleak

import (
	"testing"
)

func BenchmarkKindofLeakingCausedBySubstring(b *testing.B) {
	// var m runtime.MemStats
	for n := 0; n < b.N; n++ {
		CausedBySubstring()
	}
	// runtime.ReadMemStats(&m)
	// fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d\n", m.Alloc, m.TotalAlloc, m.Frees, m.HeapSys, m.HeapAlloc, m.HeapIdle, m.HeapReleased, m.NumGC, m.NumForcedGC)
}
