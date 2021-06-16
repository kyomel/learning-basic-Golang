package main

import "fmt"

func main() {
	name := "Michael"

	switch name {
	case "Michael":
		fmt.Println("Hello Michael")
		fmt.Println("Hello Michael")
	case "Stevan":
		fmt.Println("Hello Stevan")
		fmt.Println("Hello Stevan")
	default:
		fmt.Println("Hi, boleh kenalan?")
		fmt.Println("Hi, boleh kenalan?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
