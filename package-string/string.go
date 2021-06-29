package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Michael Stevan Lapandio", "Michael"))
	fmt.Println(strings.Split("Michael Stevan Lapandio", " "))
	fmt.Println(strings.ToLower("Michael Stevan Lapandio"))
	fmt.Println(strings.ToUpper("michael stevan lapandio"))
	fmt.Println(strings.ToTitle("michael stevan lapandio"))
	fmt.Println(strings.Trim("     Michael Stevan     ", " "))
	fmt.Println(strings.ReplaceAll("Stevan Stevan Stevan Stevan", "Stevan", "Agus"))
}
