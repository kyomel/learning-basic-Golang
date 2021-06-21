package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var michael Customer
	michael.Name = "Michael"
	michael.Address = "Poso"
	michael.Age = 27

	fmt.Println(michael.Name)
	fmt.Println(michael.Address)
	fmt.Println(michael.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Solo",
		Age:     30,
	}
	fmt.Println(joko)

	// rentan error untuk step yang dibawah bergantung pada posisi deklarasi struct
	budi := Customer{"Budi", "Jakarta", 33}
	fmt.Println(budi)

}
