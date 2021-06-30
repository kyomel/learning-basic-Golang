package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Michael")
	data.PushBack("Stevan")
	data.PushBack("Lapandio")
	data.PushFront("Joko")

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	// dimulai dari depan ke belakang
	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// dimulai dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
