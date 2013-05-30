package main

import (
	"fmt"
	"test/sem" // Implementierung von S. 171
)

func main() {
	s := sem.New(0)
	s.P()
	fmt.Println("Hello World!\n")
	s.V()
}
