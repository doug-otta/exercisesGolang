package main

import "fmt"

const helloEnglish = "Hello, "
const helloPortuguese = "Olá, "
const helloFrech = "Bonjour, "

func greetings(name, language string) string {
	switch language {
	case "portuguese":
		return helloPortuguese + name
	case "french":
		return helloFrech + name
	case "english":
		return helloEnglish + name
	default:
		return helloEnglish + "World"
	}
}

func main() {
	fmt.Println(greetings("", ""))
}
