package main

import "fmt"

// Defer function Example
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication(value int) {
// 	defer logging() //defer function
// 	fmt.Println("Run application")
// 	result := 10 / value
// 	fmt.Println("Result", result)
// }

func endApp() {
	message := recover() // recover function
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR") //panic function
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApplication(10)
	runApp(true)
	fmt.Println("Michael") // This function check about your recover is works or nah
}
