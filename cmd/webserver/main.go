package main

import (
	"fmt"
	"log"
	"net/http"

	obstrat "github.com/rgamel/obstrat-cli"
)

func main() {
	store, close, err := obstrat.FileSystemStratStoreFromFile("../../strats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := obstrat.NewStratServer(store)
	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
