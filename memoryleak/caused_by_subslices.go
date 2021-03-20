package memoryleak

var data0 []byte

func CausedBySubslices() {
	data := allocSliceMemory(1 << 20) // 1MB

	gn(data)
	// gnFix(data)
}

func gn(data1 []byte) {
	data0 = data1[len(data1)-50:]
}

func gnFix(data1 []byte) {
	data0 = append([]byte(nil), data1[len(data1)-50:]...)
}
