package main

import "fmt"

type language string

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

func greet(l language) string {
	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
