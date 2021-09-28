package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanishLang = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanishLang {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
