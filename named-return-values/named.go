package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Michael"
	middleName = "Stevan"
	lastName = "Lapandio"
	// cara lama
	// return firstName, middleName, lastName
	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
