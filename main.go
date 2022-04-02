package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var nFlag = flag.Int("n", 6, "Number of guesses")
var wordlistFlag = flag.String("wordlist", "", "Use a custom wordlist file")

type Attempt struct {
	word      string
	unmatched []rune
	matched   []rune
	positions map[rune]int
}

func findUnmatched(word string, matched []rune, positions map[rune]int) []rune {
	var unmatched = []rune{}

	var letters = []rune(word)
	for _, letter := range letters {
		if !strings.ContainsRune(string(matched), letter) {
			unmatched = append(unmatched, letter)
		}
	}

	for k, v := range positions {
		if []rune(word)[v] != k {
			unmatched = append(unmatched, []rune(word)[v])
		}
	}

	return unmatched
}

func findMatched(word string) []rune {
	var matched = []rune{}
	for _, rn := range word {
		if rn != '_' {
			matched = append(matched, unicode.ToLower(rn))
		}
	}

	return matched
}

func findPositions(word string) map[rune]int {
	var positions = map[rune]int{}
	for i, rn := range word {
		if rn != '_' && unicode.IsUpper(rn) {
			positions[unicode.ToLower(rn)] = i
		}
	}

	return positions
}

func isPotentialSolution(word string, unmatched, matched []rune, positions map[rune]int) bool {
	// Reject words that contain an unmatched letter.
	for _, rn := range unmatched {
		if strings.ContainsRune(word, rn) {
			return false
		}
	}

	// Reject words that don't contain a matched letter.
	for _, rn := range matched {
		if !strings.ContainsRune(word, rn) {
			return false
		}
	}

	// Reject words that don't have letters at a known position.
	for k, v := range positions {
		if []rune(word)[v] != k {
			return false
		}
	}

	return true
}

func Solve(wordlist []string, attempts []Attempt) []string {
	var unmatched = []rune{}
	var matched = []rune{}
	var positions = map[rune]int{}

	// Combine the results of all attempts
	for _, attempt := range attempts {
		unmatched = append(unmatched, findUnmatched(attempt.word, attempt.matched, attempt.positions)...)
		matched = append(matched, attempt.matched...)
		for k, v := range attempt.positions {
			positions[k] = v
		}
	}

	var solutions = []string{}
	for _, word := range wordlist {
		word = strings.ToLower(word)
		if isPotentialSolution(word, unmatched, matched, positions) {
			solutions = append(solutions, word)
		}
	}

	return solutions
}

func main() {
	flag.Parse()

	var guesses = []string{}
	if len(flag.Args()) < *nFlag {
		guesses = flag.Args()
	} else {
		guesses = flag.Args()[:*nFlag]
	}

	wordlist := defaultWordlist

	if *wordlistFlag != "" {
		fmt.Printf("Using wordlist %s", *wordlistFlag)
		data, err := ioutil.ReadFile(*wordlistFlag)
		if err != nil {
		}
		wordlist = strings.Split(string(data), "\n")
	}

	var attempts = []Attempt{}
	for _, guess := range guesses {
		var parts = strings.Split(guess, ":")
		var word, matched, positions string
		if len(parts) == 1 {
			word = strings.ToLower(parts[0])
		} else if len(parts) == 2 {
			word = strings.ToLower(parts[0])
			matched = strings.ToLower(parts[1])
		} else {
			word = strings.ToLower(parts[0])
			matched = strings.ToLower(parts[1])
			positions = strings.ToLower(parts[2])
		}

		attempts = append(attempts, Attempt{
			word:      strings.ToLower(word),
			matched:   findMatched(strings.ToLower(matched)),
			positions: findPositions(strings.ToUpper(positions)),
		})
	}

	var solutions = Solve(wordlist, attempts)

	if len(solutions) == 1 {
		fmt.Printf("The solution is %s!\n", solutions[0])
	} else if len(solutions) == 0 {
		fmt.Println("No solution found!")
	} else {
		fmt.Println("Possible solutions:")
		for _, solution := range solutions {
			fmt.Printf("%s\n", solution)
		}
	}
}
