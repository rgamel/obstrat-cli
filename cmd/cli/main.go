package main

import (
	"fmt"
	"log"

	obstrat "github.com/rgamel/obstrat-cli"
)

func main() {
	store, close, err := obstrat.FileSystemStratStoreFromFile("../../strats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println(store.GetRandomStrat())
}
