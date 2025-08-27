package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, tenSecondTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	// selectは複数のチャネル操作を待ち受けて、最初に準備ができたものを実行する
	// time.Aftrerを含めることで、タイムアウト処理を実装して永久に待ち続けることを防止できる
	select {
	// goroutineのcloseが実行されたタイミングで case <- chan のケースが実行される
	// どちらのping()が先にcloseされるかで、速いURLを判定する
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	// ping()の処理が完了したことを通知したいだけなので、struct{}型のチャネルを使う

	// channelの作成について
	// channelをvarキーワードで宣言すると、channelのゼロ値である「nil」で初期化される
	// nilのchannelを受け渡ししようとすると、デッドロックになるため、make()で初期化する必要がある
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// 実データを持たないchannelを宣言して、そのchannelをcloseすることで、受信側のgoroutineの終了を通知している
		close(ch)
	}()

	return ch
}
