package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

func GetRandomStrat() (string, error) {
	lines, err := getStratsFromFile()
	if err != nil {
		log.Fatalf("error getting strats from file: %v\n", err)
	}

	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex], nil
}

func getStratsFromFile() ([]string, error) {
	file, err := os.Open("./strats.txt")
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}
