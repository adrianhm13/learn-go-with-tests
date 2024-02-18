package main

import "fmt"

const (
	helloEnglishPrefix = "Hello, "
	helloSpanishPrefix = "Hola, "
	helloFrenchPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name string, lng string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(lng) + name
}

func getGreetingPrefix(lng string) (prefix string) {
	switch lng {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Adrian", "Spanish"))
}
