# Clousures

> A closure normally occurs when a function is declared inside the body of another, and the inner function references local variables of the outer function. At runtime, when the outer function is executed, then a closure is formed, which consists of the code of the inner function and references to any variables in the scope of the outer function that the closure needs.

_Definition by Wikipedia_

> Uma closure ocorre normalmente quando uma função é declarada dentro do corpo de outra, e a função interior referencia variáveis locais da função exterior. Em tempo de execução, quando a função exterior é executada, então uma closure é formada, que consiste do código da função interior e referências para quaisquer variáveis no escopo da função exterior que a closure necessita.

_Definição da Wikipédia_

To see more about Go's Closures click [here](https://tour.golang.org/moretypes/25).

When working with clousures in Go, we can make some mistakes that are easy to make and difficult do debug. I will show you some examples and explain how to avoid them.

## Common mistakes

### `for` loops

New variables declared in `for` loops are passed by reference. This means that we are using the same variable, but the value stored is being updated with each iteration.

Let's look at the code below and see what will happen:

```Go
package main

import "fmt"

func main() {
    var myFunctions []func()

    for i := 0; i < 3; i++ {
        myFunctions = append(myFunctions, func() {
            fmt.Println(i)
        })
    }

    for _, fn := range myFunctions {
        fn()
    }
}
```

Output:

```Go
3
3
3
```

Click [here](https://play.golang.org/p/V1ybQlmpJsl) to run the code above.

Note that `i` is declared inside of `for` loop, and it is being updated with each iteration. When we finally out of `for` loop and call all of our functions in the `myFunctions` slice, they are all referencing the same `i` variable which was set to `3` in the last iteration of `for` loop.

This same thing happens when we use the keyword `range`:

```Go
package main

import "fmt"

func main() {
	var myFunctions []func()

	for _, i := range []int{1, 2, 3} {
		myFunctions = append(myFunctions, func() {
			fmt.Println(i)
		})
	}

	for _, fn := range myFunctions {
		fn()
	}
}
```

Output:

```Go
3
3
3
```

Click [here](https://play.golang.org/p/hExRbwivtfN) to run the code above.

This is caused by the same issue as before.


To avoid this, we could create a function `doSomething(i)` that will receive the value of `i`. Because the parameters of a function in Go are [passed by value](https://golang.org/doc/faq#pass_by_value). This means that, at each iteration of the `for` loop, the function `doSomething(i)` will receive the value of `i` at that specific moment, and not a reference to the variable `i`. For example:

```Go
package main

import "fmt"

func main() {
	var functions []func()

	doSomething := func(i int) func() {
		return func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		functions = append(functions, doSomething(i))
	}

	for _, fn := range functions {
		fn()
	}
}
```

Output:

```Go
0
1
2
```

Click [here](https://play.golang.org/p/sk85aGdfwGT) to run the code above.


Let's look at the code bellow and see what happens:

```Go
package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	myTime *time.Time
}

func main() {
	times := make([]MyStruct, 3)
	var now time.Time
	for _, t := range times {
		now = time.Now()
		t.myTime = &now
		time.Sleep(1 * time.Second)
	}

	for i, t := range times {
		fmt.Printf("[%d] time: %v.\n", i, t.myTime.Format(time.RFC3339))
	}
}
```

Output:

```Go
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4a39fd]

goroutine 1 [running]:
main.main()
	/tmp/sandbox131697710/prog.go:22 +0x7d
```

Click [here](https://play.golang.org/p/7AHnKmiS5wJ) to run the code above.

Oh, no! The code was broken.

Note that `t` is declared inside of `for` loop, and it is being updated with each iteration. When we finally out of `for` loop and print all `mytimes` of our structures in the `times` slice, they are all referencing the same `t` variable which was set to `nil` in the last iteration of `for` loop.


To fix this crash, we must update the `i` position value of the `times` slice instead of `t`:

```Go
package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	myTime *time.Time
}

func main() {
	times := make([]MyStruct, 3)
	var now time.Time
	for i, _ := range times {
		now = time.Now()
		times[i] = &now
		time.Sleep(1 * time.Second)
	}

	for i, t := range times {
		fmt.Printf("[%d] time: %v.\n", i, t.myTime.Format(time.RFC3339))
	}
}
```

Output:

```Go
[0] time: 2009-11-10T23:00:02Z.
[1] time: 2009-11-10T23:00:02Z.
[2] time: 2009-11-10T23:00:02Z.
```

Click [here](https://play.golang.org/p/2E3F9XPsJbK) to run the code above.

The crash was fixed successfully. We had another problem, though. The time was the same for all positions.

This is caused by the same issue as before. The `now` variable value is being updated with each iterations. When we finally out of `for` loop, they are all referencing the same `now` variable which was set to `2009-11-10T23:00:02Z` in the last iteration of `for` loop.


We can solve this problem by creating a new variable `now` and assigning it with the value of `time.Now()`:

```Go
package main

import (
	"fmt"
	"time"
)

type MyStruct struct {
	myTime *time.Time
}

func main() {
	times := make([]MyStruct, 3)
	for i, _ := range times {
		now := time.Now()
		times[i] = &now
		time.Sleep(1 * time.Second)
	}

	for i, t := range times {
		fmt.Printf("[%d] time: %v.\n", i, t.myTime.Format(time.RFC3339))
	}
}
```

Output:

```Go
[0] time: 2009-11-10T23:00:00Z.
[1] time: 2009-11-10T23:00:01Z.
[2] time: 2009-11-10T23:00:02Z.
```

Click [here](https://play.golang.org/p/LyzUlwZDngm) to run the code above.

As you can see, this is one of those mistakes that are easy to make and difficult to debug. By the way, in the last example, we use pointer. Click [here](/pointers/README.md) to see more about Go's Pointers.
