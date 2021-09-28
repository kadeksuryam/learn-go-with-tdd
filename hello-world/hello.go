package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanishLang = "Spanish"
const frenchLang = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanishLang {
		return spanishHelloPrefix + name
	}
	if language == frenchLang {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
