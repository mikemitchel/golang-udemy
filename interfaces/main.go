package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// if you don't refer to passed in value, can omit variable
	// and only define the type passed in
	// custom logic for this example
	return "hello!"
}

func (spanishBot) getGreeting() string {
	// pretend this is custom spanish logic
	return "hola!"
}
