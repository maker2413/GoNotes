package main

import "fmt"

// Constants can be declared in a block
const (
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "
)

// In golang public functions begin with a capital letter and private functions
// start with a lowercase letter.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

// prefix in this case is a named return value.
// This creates a variable called prefix in our function.
// By default this variable will default to "zero" value (ints are 0 and strings are "").
// You can return the value of this variable at any point by just calling return
// Also note that this function begins with a lowercase letter so it is private.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
