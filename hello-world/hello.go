package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name != "" {
		if language == "English" {
			return englishHelloPrefix + name
		} else if language == "Spanish" {
			return spanishHelloPrefix + name
		}
	}
	return englishHelloPrefix + "World"
}

func main() {
	fmt.Println(Hello("World", "English"))
}
