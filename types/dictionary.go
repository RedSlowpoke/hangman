package types

import (
	"math/rand"
	"time"
)

func GetRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	words := []string{"GO", "RUN", "MAIN"}
	return words[rand.Intn(len(words))]
}
