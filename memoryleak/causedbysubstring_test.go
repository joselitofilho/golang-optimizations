package memoryleak

import (
	"testing"
)

func BenchmarkKindofLeakingCausedBySubstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CausedBySubstring()
	}
}
