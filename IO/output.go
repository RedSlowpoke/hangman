package io

import (
	"fmt"
	"hangman/v1/types"
)

const (
	letterQuestion  = "Guess a letter:"
	congratulations = "You have won!"
	consoledence    = "Better luck next time\nThe word was %s"
)

func AskLetter() {
	fmt.Println(letterQuestion)
}

func ShowTheWord(g types.Game) {
	fmt.Println(g.Word())
}

func ShowTheHiddenWord(g types.Game) {
	for _, r := range g.Word() {
		if g.IsGuessedRune(r) {
			fmt.Print(string(r))
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}

func ErrorDescription(e error) {
	fmt.Println(e)
}

func Congratulate(g types.Game) {
	fmt.Println(congratulations)
}
func Consoledence(g types.Game) {
	s := fmt.Sprintf(consoledence, g.Word())
	fmt.Println(s)
}
