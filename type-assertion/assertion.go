package main

import "fmt"

func random() interface{} {
	return "Michael"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// 	resultInt := result.(int) //panic
	// 	fmt.Println(resultInt)

	// using switch case
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Uknown type")
	}
}
