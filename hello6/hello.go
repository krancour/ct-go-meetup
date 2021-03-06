package main

import (
	"fmt"
	"time"
)

func generateGreetings(greetingCh chan string) {
	greetings := []string{"Hello", "Hola", "Bonjour"}
	for _, greeting := range greetings {
		greetingCh <- greeting
		time.Sleep(time.Second)
	}
	close(greetingCh)
}

func printGreetings(greetingCh chan string) {
	for greeting := range greetingCh {
		fmt.Println(greeting)
	}
	fmt.Println("No more greetings.")
}

func main() {
	greetingCh := make(chan string)
	go generateGreetings(greetingCh)
	go printGreetings(greetingCh)
	time.Sleep(time.Second * 5)
}
