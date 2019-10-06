package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Hello from `fmt.Printf()`")
	println("Hello from `println()`")
	log.Println("Hello from `log.Println`")
	log.Fatalf("Hello from `log.Fatalf()`")
}
