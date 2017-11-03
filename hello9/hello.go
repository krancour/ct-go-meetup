package main

import (
	"fmt"
	"time"
)

func generateGreetings(greetingCh chan string) {
	greetings := []string{"Hello", "Hola", "Bonjour"}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for _, greeting := range greetings {
		<-ticker.C
		greetingCh <- greeting
	}
	close(greetingCh)
}

func printGreetings(greetingCh chan string, doneCh chan bool) {
	for greeting := range greetingCh {
		fmt.Println(greeting)
	}
	close(doneCh)
}

func main() {
	greetingCh := make(chan string)
	doneCh := make(chan bool)
	go generateGreetings(greetingCh)
	go printGreetings(greetingCh, doneCh)
	<-doneCh
}
