package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printGreetings(eb)
	printGreetings(sb)
}

func (englishbot) getGreetings() string {
	return "hello"
}

func (spanishbot) getGreetings() string {
	return "hola"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
