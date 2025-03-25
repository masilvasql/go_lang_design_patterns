package main

import "github.com/masilvsql/observerComCanais/subject"

func main() {
	store := subject.Store{}

	ch1 := make(chan string)
	ch2 := make(chan string)

	store.Register(ch1)
	store.Register(ch2)

	go func() {
		for {
			promotion := <-ch1
			println("Channel 1 received promotion: " + promotion)
		}
	}()

	go func() {
		for {
			promotion := <-ch2
			println("Channel 2 received promotion: " + promotion)
		}
	}()

	store.SetPromotion("Black Friday 50% off")
	store.SetPromotion("Cyber Monday 60% off")

	store.Remove(ch1)

	store.SetPromotion("New Year 70% off")

}
