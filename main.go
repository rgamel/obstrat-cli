package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func readAtRandom(file *os.File) (string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	randomIndex := rand.Intn(len(lines))

	return lines[randomIndex], nil
}

func main() {
	file, err := os.Open("./strats.txt")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	line, err := readAtRandom(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(line)
}
