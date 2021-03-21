package underlyingmemory

import "fmt"

func UnderlyingMemoryBlockSlice1() {
	slice := []byte{1, 2, 3}
	subslice := slice[:1]
	subslice = append(subslice, 4)
	fmt.Println("subslice:", subslice)
	fmt.Println("   slice:", slice)
}

func UnderlyingMemoryBlockSlice2() {
	slice := []byte{1, 2, 3}
	slice = slice[0:2]

	sliceA := append(slice, 4)
	fmt.Println("slice:", slice)
	sliceB := append(slice, 5)
	fmt.Println("slice:", sliceB)
	fmt.Println("slice:", sliceA)
}
