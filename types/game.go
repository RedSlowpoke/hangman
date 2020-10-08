package types

type Game struct {
	word            string
	guessedRunes    []rune
	maxRounds       int
	completedRounds int
	failedRounds    int
	gameState       GameState
}

var gameRounds = 5

func NewGame(word string) Game {
	return Game{
		word:            word,
		guessedRunes:    []rune{},
		maxRounds:       gameRounds,
		completedRounds: 0,
		failedRounds:    0,
		gameState:       GameInProgress,
	}
}

func (g Game) Word() string {
	return g.word
}

func (g Game) IsGuessedRune(r rune) bool {
	for _, rr := range g.guessedRunes {
		if r == rr {
			return true
		}
	}
	return false
}

func (g *Game) IsRuneInWord(r rune) bool {
	for _, rr := range g.Word() {
		if r == rr {
			return true
		}
	}
	return false
}

func (g *Game) AddRune(r rune) {
	g.guessedRunes = append(g.guessedRunes, r)
}

func (st *Game) CompleteSuccessfulRound() {
	st.completedRounds += 1
}
func (st *Game) CompleteUnsuccessfulRound() {
	st.completedRounds += 1
	st.failedRounds += 1
}

type GameState int

const (
	GameInProgress = iota
	GameLost
	GameWon
)

func (st *Game) WinGame() {
	st.gameState = GameWon
}

func (st *Game) LoseGame() {
	st.gameState = GameLost
}

func (st Game) IsGameWon() bool {
	for _, r := range st.word {
		if !st.IsGuessedRune(r) {
			return false
		}
	}
	return true
}

func (st Game) IsGameLost() bool {
	return st.failedRounds >= st.maxRounds
}
