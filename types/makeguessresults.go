package types

type MakeGuessResult struct {
	resultCode MakeGuessResultCode
}

type MakeGuessResultCode int

const (
	ResultGuessSuccessful = iota
	ResultGuessUnsuccessful
	ResultGuessAlreadyWas
	ResultGuessErr
)

func NewMakeGuessResult(code MakeGuessResultCode) MakeGuessResult {
	return MakeGuessResult{
		resultCode: code,
	}
}

func (r MakeGuessResult) Code() MakeGuessResultCode {
	return r.resultCode
}

func (r MakeGuessResult) IsSuccess() bool {
	if r.resultCode == ResultGuessSuccessful {
		return true
	}
	return false
}

func (r MakeGuessResult) IsFail() bool {
	if r.resultCode == ResultGuessUnsuccessful {
		return true
	}
	return false
}
