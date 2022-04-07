package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var nFlag = flag.Int("n", 6, "Number of guesses")
var debugFlag = flag.Bool("debug", false, "Enable debug output")
var wordlistFlag = flag.String("wordlist", "", "Use a custom wordlist file")

type Attempt struct {
	guess   []rune
	matched []rune
}

func findUnmatchedLetters(word, matched []rune) map[rune]bool {
	var letters = map[rune]bool{}
	for i := 0; i < len(word) && i < len(matched); i++ {
		if matched[i] != '-' && matched[i] != '+' {
			letters[unicode.ToLower(word[i])] = true
		}
	}

	return letters
}

func findMatchedLetters(word, matched []rune) map[rune]bool {
	var letters = map[rune]bool{}
	for i := 0; i < len(word) && i < len(matched); i++ {
		if matched[i] == '-' || matched[i] == '+' {
			letters[unicode.ToLower(word[i])] = true
		}
	}

	return letters
}

func findMatchedPositions(word, matched []rune) map[rune][]int {
	var positions = map[rune][]int{}
	for i := 0; i < len(word) && i < len(matched); i++ {
		if matched[i] == '+' {
			var rn = unicode.ToLower(word[i])
			positions[rn] = append(positions[rn], i)
		}
	}

	return positions
}

func findUnmatchedPositions(word, matched []rune) map[rune][]int {
	var positions = map[rune][]int{}
	for i := 0; i < len(word) && i < len(matched); i++ {
		if matched[i] != '+' {
			var rn = unicode.ToLower(word[i])
			positions[rn] = append(positions[rn], i)
		}
	}

	return positions
}

func isPotentialSolution(word string, matched, unmatched map[rune]bool, matchedPositions, unmatchedPositions map[rune][]int) bool {
	// Reject words that contain an unmatched letter.
	for rn := range unmatched {
		if strings.ContainsRune(word, rn) && matched[rn] != true {
			return false
		}
	}

	// Reject words that don't contain a matched letter.
	for rn := range matched {
		if !strings.ContainsRune(word, rn) {
			return false
		}
	}

	// Reject words that don't have letters at a known position.
	for rn, positions := range matchedPositions {
		for _, i := range positions {
			if []rune(word)[i] != rn {
				return false
			}
		}
	}

	// Reject words that have letters at positions they are known _not_ to be.
	for rn, positions := range unmatchedPositions {
		for _, i := range positions {
			if []rune(word)[i] == rn {
				return false
			}
		}
	}

	return true
}

func Solve(wordlist []string, attempts []Attempt) []string {
	var matchedLetters = map[rune]bool{}
	var unmatchedLetters = map[rune]bool{}
	var matchedPositions = map[rune][]int{}
	var unmatchedPositions = map[rune][]int{}

	// Combine the results of all attempts
	for _, attempt := range attempts {
		for k, v := range findMatchedLetters(attempt.guess, attempt.matched) {
			matchedLetters[k] = v
		}
		for k, v := range findUnmatchedLetters(attempt.guess, attempt.matched) {
			unmatchedLetters[k] = v
		}
		for k, v := range findMatchedPositions(attempt.guess, attempt.matched) {
			matchedPositions[k] = append(matchedPositions[k], v...)
		}
		for k, v := range findUnmatchedPositions(attempt.guess, attempt.matched) {
			unmatchedPositions[k] = append(unmatchedPositions[k], v...)
		}
	}

	if *debugFlag {
		fmt.Printf("matched:")
		for k := range matchedLetters {
			fmt.Printf(" %c", k)
		}
		fmt.Println()
		fmt.Printf("unmatched:")
		for k := range unmatchedLetters {
			fmt.Printf(" %c", k)
		}
		fmt.Println()
		fmt.Printf("positions:")
		for k, v := range matchedPositions {
			fmt.Printf(" %c:%d", k, v)
		}
		fmt.Println()
		fmt.Printf("unpositions:")
		for k, v := range unmatchedPositions {
			fmt.Printf(" %c:%d", k, v)
		}
		fmt.Println()
	}

	var solutions = []string{}
	for _, word := range wordlist {
		word = strings.ToLower(word)
		if isPotentialSolution(word, matchedLetters, unmatchedLetters, matchedPositions, unmatchedPositions) {
			solutions = append(solutions, word)
		}
	}

	return solutions
}

func main() {
	flag.Parse()

	var args = []string{}
	if len(flag.Args()) < *nFlag {
		args = flag.Args()
	} else {
		args = flag.Args()[:*nFlag]
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
	for _, arg := range args {
		var parts = strings.Split(arg, ":")
		var guess, matched string
		if len(parts) == 1 {
			guess = strings.ToLower(parts[0])
		} else {
			guess = strings.ToLower(parts[0])
			matched = strings.ToLower(parts[1])
		}

		attempts = append(attempts, Attempt{
			guess:   []rune(guess),
			matched: []rune(matched),
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
