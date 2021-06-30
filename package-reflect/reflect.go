package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
		// old school write
		// if field.Tag.Get("required") == "true" {
		// 	value := reflect.ValueOf(data).Field(i).Interface()
		// 	if value == "" {
		// 		return false
		// 	}
		// }
	}
	return true
}

func main() {
	sample := Sample{"Michael"}

	// old scholl write
	// var sampleType reflect.Type = reflect.TypeOf(sample)
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // reflect tagnya tidak ada maka pas di run akan kosong atau tidak ditampilkan

	// sample.Name = ""
	fmt.Println(IsValid(sample))

	// hasil akan selau true karena reflect tag required tidak dicantumkan
	contoh := ContohLagi{"", ""}
	fmt.Println(IsValid(contoh))
}
