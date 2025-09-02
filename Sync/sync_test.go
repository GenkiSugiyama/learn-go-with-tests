package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// WaitGroupは、ゴルーチンのコレクションが完了するのを待つ
		var wg sync.WaitGroup
		// メインのゴルーチンはAddを呼び出して、待機するゴルーチンの数を設定する
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()

				// 各ゴルーチンが実行され、完了したらDoneを呼び出して、待機するゴルーチンの数をデクリメントする
				wg.Done()
			}(&wg)
		}

		// すべてのゴルーチンが完了するまで、Waitを使用してブロックすることができる
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

// sync.Mutexをフィールドにもつ構造体をそのまま渡してしまうと、sync.Mutexのコピーが発生してしまい、正しく動作しない
// sync.Mutexを持つ構造体を別の関数に渡す場合はポインタ型で渡す
func assertCounter(t *testing.T, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
