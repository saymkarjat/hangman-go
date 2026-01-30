package game

import (
	"fmt"
	"hangman-go/words"
	"slices"
)

type Game struct {
	playerName     string
	isOver         bool
	attemptsCount  int
	guessedLetters []rune
}

func New(playerName string) *Game {
	return &Game{
		playerName: playerName,
	}
}

func (g *Game) incrementAttemptsCount() {
	g.attemptsCount++
}

func (g *Game) Play(playerName string) {
	g.startGame(playerName)
	randomWord := g.chooseDifficulty().RandomWord()
	encodedWord := g.maskOriginalWord(randomWord)
	fmt.Println(encodedWord)
	for !g.isOver {
		g.inputLetter()
		encodedWord := g.maskOriginalWord(randomWord)
		fmt.Println(encodedWord)
		if g.attemptsCount == 6 {
			g.stopGame()
		}
	}
}

func (g *Game) startGame(playerName string) {
	fmt.Printf("Привет уважаемый " + playerName +
		", введи сложность игры(easy, medium, hard): ")
}

func (g *Game) chooseDifficulty() words.WordList {
	for {
		difficulty := inputString()
		wordList, err := words.GetWordList(
			words.Difficulty(difficulty),
		)
		if err == nil {
			return wordList
		}
		fmt.Printf("ошибка, введите сложность " +
			"игры(easy, medium, hard): ")
	}
}

func (g *Game) maskOriginalWord(word string) string {
	originalWordSlice := []rune(word)
	for i := 0; i < len(originalWordSlice); i++ {
		if !slices.Contains(
			g.guessedLetters, originalWordSlice[i],
		) {
			originalWordSlice[i] = '*'
		}
	}
	return string(originalWordSlice)
}

func (g *Game) inputLetter() {
	letter := inputString()
	g.incrementAttemptsCount()
	letters := []rune(letter)
	if (len(letters) > 1) || (len(letters) == 0) {
		return
	}
	g.guessedLetters = append(g.guessedLetters, letters...)
}

func (g *Game) stopGame() {
	g.isOver = true
	println("вы проиграли")
}

func inputString() string {
	var input string
	_, _ = fmt.Scanln(&input)
	return input
}
