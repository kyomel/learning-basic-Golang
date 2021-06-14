package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpMichael NoKTP = "1111111111"
	var marriedStatus Married = false
	fmt.Println(ktpMichael)
	fmt.Println(marriedStatus)
	fmt.Println(NoKTP("2222222222"))
}
