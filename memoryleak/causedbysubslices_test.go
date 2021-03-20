package memoryleak

import (
	"testing"
)

func BenchmarkKindofLeakingCausedBySubslices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CausedBySubslices()
	}
}
