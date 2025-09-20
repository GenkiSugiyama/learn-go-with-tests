package poker

// 1試合分の開始から終了までの責務をもつ
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
