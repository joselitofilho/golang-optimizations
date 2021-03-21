package memoryleak

var pointer *int
var number0 int

type myStruct struct {
	data   [(1 << 20)]int // 1MB
	number int
}

func CausedByInteriorPointer() {
	hn()
	// hnFix()
}

func hn() {
	ms := myStruct{number: 1}
	pointer = &ms.number
}

func hnFix() {
	ms := myStruct{number: 1}
	number0 = ms.number
}
