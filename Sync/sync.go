package sync

import "sync"

type Counter struct {
	// sync.Mutexは排他制御のための仕組み
	// 複数のgoroutineが同時に単一のリソースにアクセスするのを防ぐ
	mc    sync.Mutex
	value int
}

func NewCounter() *Counter {
	// Mutexのコピーを防ぐためにポインタ型を返す
	return &Counter{}
}

func (c *Counter) Inc() {
	// Valueの先頭でLock()でロックを取得
	c.mc.Lock()

	// deferを使って、関数が終了する直前にUnlock()を呼び出してロックを解放する
	defer c.mc.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
