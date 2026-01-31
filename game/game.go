package game

import (
	"fmt"
	"hangman-go/words"
	"slices"
	"strings"
)

type Game struct {
	playerName     string
	isOver         bool
	failedAttempts int
	guessedLetters []rune
}

func New(playerName string) *Game {
	return &Game{
		playerName: playerName,
	}
}

func (g *Game) Play() {
	g.startGame()
	randomWord := g.chooseDifficulty().RandomWord()
	maskedWord := g.maskOriginalWord(randomWord)
	fmt.Println(maskedWord)
	randomWordToSlice := []rune(randomWord)
	for !g.isOver {
		g.guessLetter(randomWordToSlice)
		maskedWord := g.maskOriginalWord(randomWord)
		fmt.Println(maskedWord)
		g.checkGameOver(maskedWord)
	}
}

func (g *Game) startGame() {
	fmt.Printf("Привет уважаемый " + g.playerName +
		", введите сложность игры(easy, medium, hard): ")
}

func (g *Game) chooseDifficulty() words.WordList {
	for {
		difficulty := scanNextLine()
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

func (g *Game) guessLetter(word []rune) {
	letter := scanNextLine()
	letters := []rune(letter)
	if (len(letters) > 1) || (len(letters) == 0) {
		fmt.Println("введите ровно одну букву")
		return
	}
	g.guessedLetters = append(g.guessedLetters, letters...)
	if !slices.Contains(word, letters[0]) {
		g.incrementFailedAttempts()
	}
	g.printHangmanStages(g.failedAttempts)
}

func (g *Game) stopGame() {
	g.isOver = true
	fmt.Println("вы проиграли")
}

func (g *Game) checkGameOver(maskedWord string) {
	if !strings.Contains(maskedWord, "*") {
		g.isOver = true
		fmt.Println("победа")
	}
	if g.failedAttempts >= 6 {
		g.stopGame()
	}
}

func (g *Game) incrementFailedAttempts() {
	g.failedAttempts++
}

func (g *Game) printHangmanStages(index int) {
	fmt.Println(hangmanStages[index])
}
func scanNextLine() string {
	var input string
	_, _ = fmt.Scanln(&input)
	return input
}
