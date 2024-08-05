package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	mizo    = "Mizo"

	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	mizoHelloPrefix    = "Chibai "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case mizo:
		prefix = mizoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
