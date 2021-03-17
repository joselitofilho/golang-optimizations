# Memory Leaking Scenarios

Generally, when programming in a language that supports automatic garbage collection, we do not have to worry about memory leaking problems, as regular cleaning of unused memory will be performed. However, we must keep in mind some special scenarios that can cause a memory leak. Next, we'll look at some scenarios.

To understand about memory management in Go and the garbage collector see [here](https://deepu.tech/memory-management-in-golang/).

## Caused by Substrings

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
