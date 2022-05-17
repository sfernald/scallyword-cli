package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	numGuesses = 3
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// allow each letter to be picked multiple times
func pickRandomLettersWithDups(sIn string, count int) string {
	s := make([]int, count, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		randNum := r.Intn(len(sIn) - 1)
		s[i] = randNum
	}

	charsIn := []rune(sIn)
	return string(charsIn[s[0]]) + string(charsIn[s[1]]) + string(charsIn[s[2]]) + string(charsIn[s[3]])
}

func main() {
	sourceWord := ""
	scallyWord := ""
	searchCount := 0
	foundWord := false
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// read in six letter word file
	// pick a random line and then a random word
	// load in 4 letter word list
	// randomly pick four letters from six letter word and do search for valid 4 letter word
	// once you have that, the game can begin
	valid4Words, err1 := os.ReadFile("four-letter-words.txt")
	check(err1)

	valid6Words, err2 := os.ReadFile("six-letter-words.txt")
	check(err2)

	word4List := strings.Fields(string(valid4Words))
	word6List := strings.Fields(string(valid6Words))

	// select random scallyword
	for foundWord == false {
		// randomly pick four letters of six letter word
		searchCount++
		sourceWord = word6List[r.Intn(len(word6List))-1]
		scallyWord = pickRandomLettersWithDups(sourceWord, 4)

		for _, word := range word4List {
			if word == scallyWord {
				foundWord = true
			}
		}
	}

	lettersCorrect := 0
	response := ""
	guess := ""
	guesses := 0
	validGuess := false

	fmt.Println("The word today is " + sourceWord)
	fmt.Println("What is the Scallyword?")

	for guesses < numGuesses {
		validGuess = false
		response = ""
		for validGuess == false {
			var reader = bufio.NewReader(os.Stdin)
			guess, _ = reader.ReadString('\n')
			guess = strings.TrimSpace(guess)

			for _, word := range word4List {
				if word == guess {
					validGuess = true
					break
				}
			}

			if validGuess == false {
				fmt.Println("Please use a real four letter word.")
			}
		}

		charsAnswer := []rune(scallyWord)
		charsGuess := []rune(guess)
		for i, _ := range charsGuess {
			if charsGuess[i] == charsAnswer[i] {
				response = response + string(charsGuess[i]) + "*"
				lettersCorrect++
			} else {
				response = response + string(charsGuess[i]) + " "
			}
		}

		guesses++
		if guess == scallyWord {
			fmt.Println("Hey, you got the scallyword!")
			break
		} else {
			fmt.Println("Sorry, wrong guess!")
			fmt.Println(response)

			if guesses == numGuesses {
				fmt.Println("Sorry, out of guesses! You lose.")
				fmt.Println("The Scallyword was " + scallyWord + ".")
			}
		}
	}
}
