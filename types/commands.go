package types

import (
	"fmt"
	"strings"
)

// list of commands that maybe given in input
type Command struct {
	code CommandCode
}

type CommandCode int

const (
	Start CommandCode = iota
	Hint
)

var Alliasses = map[string]CommandCode{
	"START": Start,
	"HINT":  Hint,
}

func AlliasToCommand(s string) ([]Command, error) {
	for a, c := range Alliasses {
		if strings.EqualFold(a, s) {
			return []Command{Command{code: c}}, nil
		}
	}
	return []Command{}, fmt.Errorf("%s is not allias for any Command", s)
}

func (s CommandCode) String() string {
	return [...]string{"Invalid command",
		"START",
		"HINT"}[s]
}
