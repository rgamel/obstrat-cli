package obstrat

import (
	"bufio"
	"math/rand"
	"os"
)

func GetRandomStrat(lines []string) string {
	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex]
}

func GetStratsFromFile(file *os.File) ([]string, error) {
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

func ShuffleStratPile(a []string) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func DrawTopStrat(lines []string) string {
	// simulates drawing top card,
	// reading (returning) it,
	// and placing it back on the bottom of the deck
	head := lines[0]
	lines = append(lines[1:], head)

	return head
}
