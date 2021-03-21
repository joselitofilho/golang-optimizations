# Underlying Memory Block
_see code [here](underlying_memory_block.go)_

## Arrays and Slices

### Array

Arrays are just continuous blocks of memory. We can simply represent array as a set of blocks in memory sitting next to each other.

```Go
var arr [3]
var arr [3]byte{1, 2, 3}
var arr [...]byte{1, 2, 3}
```

<img src="media/array.png" width=200>

### Slice

Slices are similar to arrays, and the declaration is really similar:

```Go
var slice []byte
```

But Go's slices ([src/runtime/slice.go](https://golang.org/src/runtime/slice.go#L11)) are structures with three fields - pointer to array, length and capacity:

```Go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

When you create a new slice, Go runtime will create a three-block object in memory with the `array` set to `nil` and `len` and `cap` set to `0`.

```Go
var slice = make([]byte, 3)
var slice = []byte{1, 2, 3}
```

<img src="media/slice.png" width=350>

That was easy. But what will happen, if you create another subslice and change some elements? Let's try:

```Go
slice := []byte{1, 2, 3}
subslice := slice[:1]
subslice = append(subslice, 4)
fmt.Println("subslice:", subslice)
fmt.Println("   slice:", slice)
```

Output:

```bash
subslice: [1 4]
   slice: [1 4 3]
```

[https://play.golang.org/p/HC43cpEKtF5](https://play.golang.org/p/HC43cpEKtF5)

![subslice](media/subslice.gif)


Similar result is with the following code:

```Go
slice := []byte{1, 2, 3}
slice = slice[0:2]

sliceA := append(slice, 4)
fmt.Println("slice:", slice)
sliceB := append(slice, 5)
fmt.Println("slice:", sliceB)
fmt.Println("slice:", sliceA)
```

Output:

```bash
slice: [1 2]
slice: [1 2 5]
slice: [1 2 5]
```

[https://play.golang.org/p/GShOrk-8Pza](https://play.golang.org/p/GShOrk-8Pza)

Again, both the `sliceA` and `sliceB` are reusing the same underlying array.

To avoid or prevent this surprise, never keep both the slice and its subslice around. That is, always do something either like `slice = slice[1:]` or make a copy.


The standard Go compiler/runtime does let them share the same underlying memory block. This is a good design, which is both memory and CPU consuming wise. But it may cause [memory leaking](/memoryleak/README.md) sometimes.
