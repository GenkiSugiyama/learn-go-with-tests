package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/GenkiSugiyama/learn-go-with-tests/http-server"
)

type GameSpy struct {
	StartCalledWith  int
	FinishCalledWith string
	StartCalled      bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalledWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

var dummyBlidAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}
var dummyGameSpy = &GameSpy{}

func TestCLI(t *testing.T) {
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		it := userSends("Pies")
		game := &GameSpy{}

		cli := poker.NewCLI(it, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)

		if game.StartCalledWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})

	t.Run("finish game with Chris as winner", func(t *testing.T) {
		in := userSends("1", "Chris wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishCalledWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishCalledWith)
		}
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := userSends("1", "Cleo wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.FinishCalledWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishCalledWith)
		}
	})
}

func userSends(inputs ...string) io.Reader {
	return strings.NewReader(strings.Join(inputs, "\n") + "\n")
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
