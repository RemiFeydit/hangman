package game

import (
	"bufio"
	"math/rand"
	"os"
	"time"
	"hangman.com/notitou/hangman/utils"
)

func RandomWord() string {
	if utils.IsParam(){
		f, _ := os.Open(os.Args[1])
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
	os.Exit(3)
	return ""
}
