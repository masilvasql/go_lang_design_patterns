package main

import (
	"fmt"
	"github.com/masilvasql/facade/home_theater"
)

func main() {
	ht := home_theater.NewHomeTheaterFacade()

	fmt.Println("🎬 Starting home theater")
	ht.StartMoviewNight()

	fmt.Println("\n🛏 Ending home theater")
	ht.EndMovieNight()
}
