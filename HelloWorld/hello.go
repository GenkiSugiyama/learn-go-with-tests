package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
