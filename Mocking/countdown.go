package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// io.WriterとSleeperインターフェースのどちらも実装するSpy構造体
// Sleep()とWrite()それぞれで”sleep”, "write"をCallスライスに追加することで、
// どの順番で呼び出されたかを確認できる。テスト用のSpyとして利用する。
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleeperインターフェースを実装するためのメソッド
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// io.Writerインターフェースを実装するためのメソッド
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// time.Sleep()やMockのSleep()を注入できるようにするための構造体
// 本番ではtimeの実装を注入したり、テストではSpyを注入したりできる
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
