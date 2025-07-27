package main

import "fmt"


const (
englishHello = "Hello "
spanish = "Spanish"
spanishHello = "Hola "
french = "French"
frenchHello = "Bonjour "
)

func main() {
	fmt.Println(Hello("World", "English"))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case french: prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}
	return
}
