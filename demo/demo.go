package main

import (
	"fmt"
	"log"
	"time"

	"github.com/krancour/ct-go-meetup/demo/meetup"
)

const delay = 1

func lineUpGophers(gopherCh chan meetup.Gopher) {
	gophers := []meetup.Gopher{
		&meetup.NoviceGopher{
			Name:    "Tony Kerz",
			Company: "Aetna",
		},
		&meetup.ExperiencedGopher{
			Name:            "Kent Rancourt",
			Company:         "Microsoft",
			YearsExperience: 3,
		},
		&meetup.ExperiencedGopher{
			Name:            "Rob Pike",
			Company:         "Google",
			YearsExperience: 9,
		},
	}
	ticker := time.NewTicker(time.Second * delay)
	defer ticker.Stop()
	for _, gopher := range gophers {
		<-ticker.C
		gopherCh <- gopher
	}
	close(gopherCh)
}

func printGreetings(gopherCh chan meetup.Gopher, errCh chan error) {
	for gopher := range gopherCh {
		greeting, err := gopher.GetGreeting()
		if err != nil {
			errCh <- err
		} else {
			fmt.Println(greeting)
		}
	}
	fmt.Println("That's all the Gophers!")
	close(errCh)
}

func main() {
	gopherCh := make(chan meetup.Gopher)
	errCh := make(chan error)
	go lineUpGophers(gopherCh)
	go printGreetings(gopherCh, errCh)
	if err := <-errCh; err != nil {
		log.Fatal(err)
	}
}
