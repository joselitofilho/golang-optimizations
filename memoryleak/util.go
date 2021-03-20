package memoryleak

func allocStringMemory(size int) string {
	return string(make([]byte, size))
}

func allocSliceMemory(size int) []byte {
	return make([]byte, size)
}
