package main

import (
	"Basic-Golang/database"   // package inizialization
	_ "Basic-Golang/database" // blank identifier
	"Basic-Golang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Michael")
	helper.SayGoodbye("Michael")
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error

	// part package inizialization
	result := database.GetDatabase()
	fmt.Println(result)
}
