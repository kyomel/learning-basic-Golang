package main

import "fmt"

func getFullName() (string, string, string) {
	return "Michael", "Stevan", "Lapandio"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	// fmt.Println(middleName)
	fmt.Println(lastName)
}
