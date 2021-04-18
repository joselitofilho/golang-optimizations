# Pointers

One of the most porwerful features in programming is pointers. They allow us to directly access objects in memory, allowing us to optimize our algorithms.

Learning how to use them can be difficult and debugging code with pointer problems is a challange.

<p align="center"><img src="media/pointers.gif"></p>

```Go
package main

import "fmt"

func main() {
	var name string = "JF"
	pname := &name         // pointer to name

	fmt.Println("Variable content: ", name)
	fmt.Println("Pointer content:  ", *pname)

	fmt.Println("------------------------------------")

	fmt.Println("Variable address: ", &name)
	fmt.Println("Pointer address:  ", pname)

	fmt.Println("------------------------------------")
	
	*pname = "Joselito"    // set name through the pointer
	
	fmt.Println("Variable content: ", name)
	fmt.Println("Pointer content:  ", *pname)
}

Output:

Variable content:  JF
Pointer content:   JF
------------------------------------
Variable address:  0xc000010240
Pointer address:   0xc000010240
------------------------------------
Variable content:  Joselito
Pointer content:   Joselito
```

Click [here](https://play.golang.org/p/v0qKfKMSMDn) to run the code above.