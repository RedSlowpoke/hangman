package types

import (
	"fmt"
)

type ReadInputError struct {
	code ReadInputErrorCode
}

type ReadInputErrorCode int

const (
	ErrInvalidInputRead ReadInputErrorCode = iota
	ErrEmptyInput
	ErrMoreThanOneCharacter
	ErrNotLetter
)

func (s ReadInputErrorCode) String() string {
	return [...]string{"Unknown error",
		"Please type one symbol or command",
		"Please type one symbol or command",
		"Please use only latin alphabet"}[s]
}

func (e *ReadInputError) Error() string {
	return fmt.Sprintf("error: %s", e.code)
}

func NewReadInputError(code ReadInputErrorCode) ReadInputError {
	return ReadInputError{
		code: code,
	}
}

type GameError struct {
	code GameErrorCode
}

type GameErrorCode int

const (
	ErrInvalidGameAction GameErrorCode = iota
	ErrRuneAlreadyGuessed
)

func (s GameErrorCode) String() string {
	return [...]string{"Unknown error",
		"Rune already was guessed this game"}[s]
}

func (e *GameError) Error() string {
	return fmt.Sprintf("error: %s", e.code)
}

func NewGameError(code GameErrorCode) GameError {
	return GameError{
		code: code,
	}
}
