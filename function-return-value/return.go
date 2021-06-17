package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Kang"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Michael Stevan")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
