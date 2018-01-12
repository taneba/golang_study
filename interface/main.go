package main

import "fmt"

// getGreeting関数をもつtypeは、bot typeでもあるようになる
// interfaceは、振る舞いの似たオブジェクトを同じものとして分類することで共通した処理をまとめる目的がある
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type japaneseBot struct{}

func main() {
	eb := englishBot{}
	jp := japaneseBot{}

	printGreeting(eb)
	printGreeting(jp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}

func (japaneseBot) getGreeting() string {
	return "やあ"
}
