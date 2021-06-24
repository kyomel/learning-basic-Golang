package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	// address2 := address1

	// address2.City = "Jakarta"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// penulisan pointer lama
	// var address1 Address = Address{"Sleman", "Yogyakarta", "Indonesia"}
	// var address2 *Address = &address1

	// pointer
	address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	address2 := &address1
	address3 := &address1
	// Membuat memory yg baru
	// address2.City = "Jakarta"
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Yogyakarta"
	fmt.Println(address4)

}
