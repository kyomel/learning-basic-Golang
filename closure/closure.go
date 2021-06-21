package main

import "fmt"

func main() {
	name := "Joko"
	counter := 0

	increment := func() {
		name := "Tukiyem"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
