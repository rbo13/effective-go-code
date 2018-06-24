package main

import "fmt"

// Bot a simple interface
// that handles greeting
type Bot interface {
	getGreeting() string
}
type englishBot struct {
}
type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Ola!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
