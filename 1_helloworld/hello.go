package main

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello World ("
	spanishHelloPrefix = "Hola ("
	frenchHelloPrefix  = "Bonjour ("

	globalHelloPostfix = ")"
)

func Hello(name string, language string) string {
	if name == "" {
		return "Hello World"
	}

	// if language == spanish {
	// 	return spanishHelloPrefix + name + globalHelloPostfix
	// }

	// prefix := englishHelloPrefix

	// switch language {
	// case spanish:
	// 	prefix = spanishHelloPrefix
	// case french:
	// 	prefix = frenchHelloPrefix
	// }

	// return prefix + name + globalHelloPostfix

	return greetingPrefix(language) + name + globalHelloPostfix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

// func main() {
// 	fmt.Println(Hello("Yay"))
// }
