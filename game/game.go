package game

import (
	"fmt"
	"hangman-go/words"
)

type Game struct {
	playerName string
	isOver     bool
}

func New(playerName string) *Game {
	return &Game{
		playerName: playerName,
	}
}

func (g *Game) Play(playerName string) {
	for !g.isOver {
		g.startGame(playerName)
	}
}

func (g *Game) startGame(playerName string) {
	fmt.Printf("Привет уважаемый " + playerName + ", введи сложность игры(easy, medium, hard): ")
	randomWord := g.chooseDifficulty().RandomWord()
	fmt.Println(randomWord)
}

func (g *Game) chooseDifficulty() words.WordList {
	for {
		difficulty := inputString()
		wordList, err := words.GetWordList(words.Difficulty(difficulty))
		if err == nil {
			return wordList
		}
		fmt.Printf("ошибка, введите сложность игры(easy, medium, hard): ")
	}
}

func inputString() string {
	var input string
	_, _ = fmt.Scanln(&input)
	return input
}
