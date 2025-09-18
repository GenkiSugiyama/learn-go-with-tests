package poker_test

import (
	"strings"
	"testing"

	poker "github.com/GenkiSugiyama/learn-go-with-tests/http-server"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user imput", func(t *testing.T) {
		// ユーザーからの入力をシミュレートする
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		// ユーザーからの入力をシミュレートする
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
