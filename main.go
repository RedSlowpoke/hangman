package main

import (
	"fmt"
	"hangman/v1/commands"
	"hangman/v1/io"
)

func main() {
	Play()
}

func Play() {
	game := commands.Start()
	io.ShowTheWord(game)
	for {
		io.AskLetter()
		io.ShowTheHiddenWord(game)
		input, err := io.ReadInput()
		if err != nil {
			io.ErrorDescription(err)
			continue
		}
		guessResult, err := commands.MakeGuess(&game, input.Rune())
		if err != nil {
			io.ErrorDescription(err)
			continue
		}
		if guessResult.IsSuccess() {
			game.CompleteSuccessfulRound()
		}
		if guessResult.IsFail() {
			game.CompleteUnsuccessfulRound()
		}
		fmt.Println(err)
		io.ShowTheHiddenWord(game)
		fmt.Println(game)
		if game.IsGameWon() {
			game.WinGame()
			io.Congratulate(game)
			return
		}
		if game.IsGameLost() {
			game.LoseGame()
			io.Concoledence(game)
			return
		}
	}
}
