package main

import (
	"Basic-Golang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Michael")
	helper.SayGoodbye("Michael")
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
