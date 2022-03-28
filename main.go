package main

import "fmt"

var version = "dev"

func main() {

	fmt.Printf("Version1: %s\n", version)

	fmt.Println(hello())
}

func hello() string {
	return "Hello Golang"
}
