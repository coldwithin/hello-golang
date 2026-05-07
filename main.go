package main

import "fmt"

func main() {
	fmt.Println(hello(), "this is some changed text")
}

func hello() string {
	return "Hello Go"
}
