package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findExist(word string) []rune {
	var exists = make([]rune, 0)
	for _, rn := range word {
		if rn != '_' {
			exists = append(exists, unicode.ToLower(rn))
		}
	}

	return exists
}

func findPosition(word string) map[rune]int {
	var positions = make(map[rune]int)
	for i, rn := range word {
		if unicode.IsUpper(rn) {
			positions[unicode.ToLower(rn)] = i
		}
	}

	return positions
}

func isPotentialSolution(word string, enteredWords []string, knownExist []rune, knownPosition map[rune]int) bool {
	// If a word was already entered it can't be a possible solution.
	for _, enteredWord := range enteredWords {
		if word == enteredWord {
			return false
		}
	}

	// Reject words that don't contain a known existing letter.
	for _, letter := range knownExist {
		if !strings.ContainsRune(word, letter) {
			return false
		}
	}

	var letters = make([]rune, len(word))
	for i, rn := range word {
		letters[i] = rn
	}

	// Reject words that don't have letters at a known position.
	for k, v := range knownPosition {
		if letters[v] != k {
			return false
		}
	}

	return true
}

func ShowSolutions(enteredWords []string, knownExist []rune, knownPosition map[rune]int) bool {
	var matches = make([]string, 0)
	for _, word := range wordlist {
		if isPotentialSolution(word, enteredWords, knownExist, knownPosition) {
			matches = append(matches, word)
		}
	}
	fmt.Print("Entered words:")
	for _, enteredWord := range enteredWords {
		fmt.Printf(" %s", enteredWord)
	}
	//fmt.Println()

	fmt.Print("Known letters:")
	for _, letter := range knownExist {
		fmt.Printf(" %c", letter)
	}
	fmt.Println()

	fmt.Print("Known Positions:")
	for k, v := range knownPosition {
		fmt.Printf(" %c:%d", k, v)
	}
	fmt.Println()

	if len(matches) == 1 {
		fmt.Printf("The solution is: %s\n", matches[0])
		return true
	}

	if len(matches) == 0 {
		fmt.Println("No solutions found!")
		return false
	}

	fmt.Println("Possible solutions:")
	for _, match := range matches {
		fmt.Printf("%s\n", match)
	}

	return false
}

func main() {
	flag.Parse()

	if flag.Arg(0) != "" {
		var word = flag.Arg(0)
		var knownExist = findExist(word)
		var knownPosition = findPosition(word)

		ShowSolutions([]string{word}, knownExist, knownPosition)
	} else {
		var reader = bufio.NewReader(os.Stdin)
		var enteredWords = make([]string, 0)
		for {
			fmt.Print("Enter a word: ")
			word, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			enteredWords = append(enteredWords, word)

			fmt.Print("Enter known letters: ")
			exists, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			fmt.Print("Enter known letters at known positions ('_' is unknown): ")
			positions, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			var knownExists = findExist(strings.TrimSuffix(exists, "\n"))
			var knownPositions = findPosition(strings.ToUpper(strings.TrimSuffix(positions, "\n")))

			fmt.Println()
			if ShowSolutions(enteredWords, knownExists, knownPositions) {
				break
			}
		}
	}
}
