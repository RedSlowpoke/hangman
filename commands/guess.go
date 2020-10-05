package commands

import (
	"hangman/v1/types"
)

func MakeGuess(g *types.Game, r rune) (types.MakeGuessResult, error) {
	if g.IsGuessedRune(r) {
		err := types.NewGameError(types.ErrRuneAlreadyGuessed)
		return types.NewMakeGuessResult(types.ResultGuessAlreadyWas), &err
	}

	if g.IsRuneInWord(r) {
		g.AddRune(r)
		return types.NewMakeGuessResult(types.ResultGuessSuccessful), nil
	} else {
		g.CompleteUnsuccessfulRound()
		return types.NewMakeGuessResult(types.ResultGuessUnsuccessful), nil
	}
}
