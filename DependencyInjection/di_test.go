package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Greet()をテストする際に、io.Writerインターフェースを実装したオブジェクトを初期化して渡す
	// 今回はbytes.Bufferを初期化して渡すことで、メモリ上に再利用可能なバッファとして利用できる
	buffer := bytes.Buffer{}
	// 標準出力はされるが、テストは失敗する
	Greet(&buffer, "Chris")

	// Greetに渡したbuffer.Bufferに書き込まれた内容を取得している
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
