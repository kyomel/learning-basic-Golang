package helper

import "fmt"

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

// function tidak dapat diekspor karena dimulai dengan huruf kecil
// func sayGoodbye(name string) {
// 	fmt.Println("Goodbye", name)
// }

func SayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
