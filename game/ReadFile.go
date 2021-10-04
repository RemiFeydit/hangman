package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(attempt int) {
	file, err := os.Open("hangman.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for index, eachline := range txtlines {
		if index >= attempt*7+attempt && index < attempt*7+7+attempt {
			fmt.Println(eachline)
		}
	}
}
