package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Print()だと標準出力されてしまいテストが失敗する
	// fmt.Printf("Hello, %s", name)
	// そのためFprintf()を使って出力先を外部から指定できるio.Writerインターフェースを受け取るようにする
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Genki")
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		// エラーが発生したらログに出力して終了する
		log.Fatal(err)
	}
}
