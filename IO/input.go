package io

import (
	"bufio"
	"hangman/v1/types"
	"os"
	"strings"
	"unicode/utf8"
)

func ReadInput() (types.ReadInputResult, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return types.NewReadInputResult(types.ResultCodeErr, []rune{}, []types.Command{}), err
		}
		// empty spaces should meaningles in input
		line = strings.TrimSpace(line)
		switch utf8.RuneCountInString(line) {
		// I hope that this should mean an empty line
		case 0:
			err := types.NewReadInputError(types.ErrEmptyInput)
			return types.NewReadInputResult(types.ResultCodeErr, []rune{}, []types.Command{}), &err
			// Should mean empty space before and after the symbol
		case 1:
			char, err := getNormalizedRuneMap(line)
			if err != nil {
				return types.NewReadInputResult(types.ResultCodeErr, []rune{}, []types.Command{}), err
			}
			return types.NewReadInputResult(types.ResultCodeChar, char, []types.Command{}), nil
			// Should mean more than one symbol, which maybe a command
		default:
			command, err := types.AlliasToCommand(line)
			if err != nil {
				return types.NewReadInputResult(types.ResultCodeErr, []rune{}, []types.Command{}), err
			}
			return types.NewReadInputResult(types.ResultCodeCommand, []rune{}, command), nil
		}
	}
}

func getNormalizedRuneMap(s string) ([]rune, error) {
	// check that this is a latin alphabet
	s = strings.ToUpper(s)
	runes := []rune(s)
	if runes[0] >= 0101 && runes[0] <= 0132 {
		return []rune{runes[0]}, nil
	}
	err := types.NewReadInputError(types.ErrNotLetter)
	return []rune{}, &err
}
