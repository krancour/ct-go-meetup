package main

import (
	"fmt"
	"sync"
	"time"
)

func generateGreetings(greetingCh chan string, wg *sync.WaitGroup) {
	greetings := []string{"Hello", "Hola", "Bonjour"}
	for _, greeting := range greetings {
		greetingCh <- greeting
		time.Sleep(time.Second)
	}
	close(greetingCh)
	wg.Done()
}

func printGreetings(greetingCh chan string, wg *sync.WaitGroup) {
	for greeting := range greetingCh {
		fmt.Println(greeting)
	}
	wg.Done()
}

func main() {
	greetingCh := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go generateGreetings(greetingCh, wg)
	go printGreetings(greetingCh, wg)
	wg.Wait()
}
