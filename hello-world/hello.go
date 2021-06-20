package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const holaPrefix = "Hola, "
const bonjourPrefix = "Bonjour, "

func greetingPrefix(language string) (prefix string) {
	// This function uses a name return value (prefix), can be returned without using name

	switch language {
	case spanish:
		prefix = holaPrefix
	case french:
		prefix = bonjourPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
