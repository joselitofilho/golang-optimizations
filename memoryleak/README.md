# Memory Leaking

_Memory leak_ is the name is given when a block of memory, allocated for a given operation, is not released when it is no longer needed. A memory leak can also happen when an object is stoted in memory but can no longer be accessed anymore by the running code.

This [post blog](https://medium.com/dm03514-tech-blog/sre-debugging-simple-memory-leaks-in-go-e0a9e6d63d4d) describes in detail the memory leak.

**Table of Contents**

- [Garbage Collector](#garbage-collector)
- [Scenarios](#scenarios)
  - [Caused by Substrings](#caused-by-substrings)
  - [Caused by Subslices](#caused-by-Subslices)
  - [Caused by Interior Pointer](#caused-by-interior-pointer)
  - [Caused by Using Finalizers Improperly](#caused-by-using-finalizers-improperly)

## Garbage Collector

Generally, when programming in a language that supports automatic garbage collection, we do not have to worry about memory leaking problems, as regular cleaning of unused memory will be performed. However, we must keep in mind some special scenarios that can cause a memory leak. Next, we'll look at some scenarios.

To understand about memory management in Go and the garbage collector see [here](https://deepu.tech/memory-management-in-golang/).

## Scenarios

### Caused by Substrings
_see code [here](caused_by_substring.go)_

The Go standard compiler/runtime allows a `s` string and an `s` substring to share the same underlying memory block. This is very good for saving memory and processing. But it can sometimes cause a memory leak.

For example, after the function `CausedBySubstring` in the following example is called, there will be about 1MB memory leaking, until the package-level variable `str0` is modified again elsewhere.

```Go
var str0 string

func CausedBySubstring() {
	str := allocMemory(1 << 20) // 1MB
	fn(str)
}

func allocMemory(size int) string {
	return string(make([]byte, size))
}

func fn(str1 string) {
	str0 = str1[:50]
}
```

![caused by substrings](media/caused-by-string.gif)

I will show you some ways to avoid this memory leaking.

The first way is to convert the substring `str1` to a `[]byte` value then convert the `[]byte` value back to `string`.

```Go
func fnFix1(str1 string) {
	str0 = string([]byte(str1[:50]))
}
```

The disadvantage of avoiding memory leaking the above way is that we duplicate the 50 bytes unnecessarily.


There is a way to avoid the unnecessary duplicate using the `strings.Builder`.

```Go
func fnFix2(str1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(str1[:50])
	str0 = b.String()
}
```

The disadvantage of the above way is it is a little verbose.


Finally, we can use the [`strings.Repeat`](https://golang.org/pkg/strings/#Repeat). Since the Go 1.12, we can call the `strings.Repeat` function with the count argument as 1 in the strings standard package to clone a string, and the implementaion of `strings.Repeat` will make use of `strings.Builder`, to avoid one unnecessary duplicate.

```Go
func fnFix3(str1 string) {
	str0 = strings.Repeat(str1[:50], 1)
}
```

### Caused by Subslices
_see code [here](caused_by_subslices.go)_

Similarly to substrings, subslices may also cause memory leaking.

```Go
var data0 []byte

func CausedBySubstring() {
	data := allocMemory(1 << 20) // 1MB
	gn(str)
}

func allocMemory(size int) []byte {
	return make([]byte, size)
}

func gn(data1 []byte) {
	str0 = str1[:50]
}
```

In the above code, after the `gn` function is called, most memory occupied by the memory block hosting the elements of `data1` will be lost.


To avoid this memory leaking, we must duplicate the 50 elements for `data0`, so that the aliveness of `data0` will not prevent the memory block hosting the elements of `data1` from being collected by the garbage collector.

```Go
func gnFix1(data1 []byte) {
	data0 = append([]byte(nil), data1[len(data1)-50:]...)
}
```

### Caused by Interior Pointer
_see code [here](caused_by_interior_pointer.go)_

After the `hn` function is called, the memory block allocated for the `myStruct` will cause memory leaking.

```Go
func hn() {
	ms := myStruct{number: 1}
	pointer = &ms.number
}
```

As long as `pointer` is still alive, it will retain `myStruct` in memory. This happens because an interior pointer to a structure's field keeps the whole struct in memory.


To avoid this memory leaking, we must make a copy of `myStruct.number`.

```Go
func hnFix() {
	ms := myStruct{number: 1}
	number0 = ms.number
}
```

### Caused by Using Finalizers Improperly
_see code [here](caused_by_finalizers.go)_

<p align="center">
<img src="media/go-finalizer.png" />
<i>Illustration created for ???A Journey With Go???, made from the original Go Gopher, created by Renee French.</i>
</p>

Go runtime provides a method [runtime.SetFinalizer](https://golang.org/pkg/runtime/#SetFinalizer) that allows to provider a function to a variable that will be called when the garbage collector sees this variable as garbage ready to be collected since it became unreachable.

After the `usingfinalizerImproperly` function is called, `x` and `y` forms a cyclic reference group. Preventing all memory blocks to the group from being collected.

```Go
func usingfinalizerImproperly() {
	type T struct {
		data [1 << 20]int // 1MB
		t    *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	runtime.SetFinalizer(&x, finalizer)

	x.t, y.t = &y, &x
}
```

So, we must avoid setting finalizers for values in a cyclic reference group.

```Go
func notUsefinalizer() {
	type T struct {
		data [1 << 20]int // 1MB
		t    *T
	}

	var x, y T

	x.t, y.t = &y, &x
}
```

By the way, I don't recommend using finalizers as object destructors.

See [here](https://play.golang.org/p/jWhRSPNvxJ) an example using finalizer.
