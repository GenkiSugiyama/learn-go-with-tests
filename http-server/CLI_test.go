package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/GenkiSugiyama/learn-go-with-tests/http-server"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

var dummyBlidAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("finish game with Chris as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
