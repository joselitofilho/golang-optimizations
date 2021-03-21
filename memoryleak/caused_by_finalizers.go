package memoryleak

import (
	"fmt"
	"runtime"
)

type T struct {
	data [1 << 20]int // 1MB
	t    *T
}

func CausedByFinalizers() {
	usingfinalizerImproperly()
	// notUsefinalizer()
}

func usingfinalizerImproperly() {
	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	runtime.SetFinalizer(&x, finalizer)

	x.t, y.t = &y, &x
}

func notUsefinalizer() {
	var x, y T

	x.t, y.t = &y, &x
}
