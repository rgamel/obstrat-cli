package main

import (
	"fmt"
	"log"
)

func main() {
	line, err := GetRandomStrat()

	if err != nil {
		log.Fatalf("error getting strat: %v\n", err)
	}

	fmt.Println(line)
}
