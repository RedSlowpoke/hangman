package commands

import (
	"hangman/v1/types"
)

func Start() types.Game {
	return (types.NewGame(types.GetRandomWord()))
}
