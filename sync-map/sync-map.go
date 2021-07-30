package main

import (
	"fmt"
	"sync"
)

type Class struct {
	Students sync.Map
}

func handler(key, value interface{}) bool {
	fmt.Printf("Name: %s - Value: %s\n", key, value)
	return true
}

func main() {
	class := &Class{}

	// Stored value
	class.Students.Store("Zhao", "class 1")
	class.Students.Store("Qian", "class 2")
	class.Students.Store("Sun", "class 3")

	// Traversing, passing in a function, when the function is traversed, the function returns false to stop traversing
	class.Students.Range(handler)

	// Inquire
	if _, ok := class.Students.Load("Li"); !ok {
		fmt.Println("--> Li not found")
	}

	// Query or append
	_, loaded := class.Students.LoadOrStore("Li", "class 4")
	if loaded {
		fmt.Println("--> Load Li success")
	} else {
		fmt.Println("--> Store Li success")
	}

	// Delete
	class.Students.Delete("Sun")
	fmt.Println("--> Delete Sun success")

	// Traversing
	class.Students.Range(handler)
}
