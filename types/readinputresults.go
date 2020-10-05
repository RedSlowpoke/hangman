package types

type ReadInputResult struct {
	resultCode ResultCode
	rune       []rune
	command    []Command
}

type ResultCode int

const (
	ResultCodeChar = iota
	ResultCodeCommand
	ResultCodeErr
)

func NewReadInputResult(code ResultCode, rune []rune, command []Command) ReadInputResult {
	return ReadInputResult{
		resultCode: code,
		rune:       rune,
		command:    command,
	}
}

func (input ReadInputResult) Rune() rune {
	return input.rune[0]
}
