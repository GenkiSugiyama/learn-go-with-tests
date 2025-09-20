package poker

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerterは、指定された時間後に指定された金額のブラインドが上がることを通知するインターフェースです。
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFuncは、BlindAlerterインターフェースを関数として実装している
type BlindAlerterFunc func(duration time.Duration, amount int)

// BlindAlerterFuncがScheduleAlertAtメソッドを実装することでBlindAlerterとして扱えるようになる
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// BlindAlerterFuncと同じシグネチャを持つためBlindAlerterFuncでラップすることで
// BlindAlerterFuncが実装しているBlindAlerterインターフェースとして扱うことができる
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
