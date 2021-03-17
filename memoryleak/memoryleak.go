package memoryleak

import "strings"

var str0 string

func CausedBySubstring() {
	str := allocMemory(1 << 20) // 1MB

	fn(str)
	// fnFix1(str)
	// fnFix2(str)
	// fnFix3(str)
}

func allocMemory(size int) string {
	return string(make([]byte, size))
}

func fn(str1 string) {
	str0 = str1[:50]
}

func fnFix1(s1 string) {
	str0 = string([]byte(s1[:50]))
}

func fnFix2(s1 string) {
	str0 = strings.Repeat(s1[:50], 1)
}

func fnFix3(s1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	str0 = b.String()
}
