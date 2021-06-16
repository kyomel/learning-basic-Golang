package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke- ", counter)
	}

	slice := []string{"Michael", "Stevan", "Lapandio", "Budi", "Joko"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice { //underscore digunakan untuk tidak menampilkan nilai dari yg ngin ditampilkan
		fmt.Println("Index", i, "=", value)
		// fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Michael Stevan"
	person["title"] = "Software Engineer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
