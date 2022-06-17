package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloSuffix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloSuffix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
