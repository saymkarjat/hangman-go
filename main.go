package main

import (
	"fmt"
	"hangman-go/game"
)

func main() {
	fmt.Printf("введите ваше имя: ")
	var playerName string
	_, _ = fmt.Scanln(&playerName)

	g := game.New(playerName)
	g.Play(playerName)
}
