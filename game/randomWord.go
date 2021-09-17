package game

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWord() string {
	f, _ := os.Open("hangman.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	wordArray := []string{}
	rand.Seed(time.Now().UnixNano())
	for scanner.Scan() {
		wordArray = append(wordArray, scanner.Text())
	}
	randWordIndex := rand.Intn(len(wordArray))
	return wordArray[randWordIndex]
}
