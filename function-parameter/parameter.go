package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	firstName := "Michael"
	sayHelloTo(firstName, "Stevan")
	sayHelloTo("Ken", "Himura")
}
