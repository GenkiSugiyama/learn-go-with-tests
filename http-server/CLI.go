package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		// io.Readerを実装するものをbufio.Scannerでラップする
		// CLIアプリなので標準入力（os.Stdin）が渡されて、標準入力をスキャンする
		in: bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	// Scanner.Text()で読み取った文字列を返す
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	// Scan()は入力から1行を読み取る, 改行まで待機する
	cli.in.Scan()
	// Text()はScan()で読み取った文字列を返す
	return cli.in.Text()
}
