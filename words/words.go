package words

import (
	"errors"
	"math/rand/v2"
)

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

type WordList []string

var (
	easyWords = WordList{
		"кот", "дом", "мяч", "сон", "лес",
	}
	mediumWords = WordList{
		"корабль", "молоток", "зеркало", "черепаха", "крокодил",
	}
	hardWords = WordList{
		"дихотомия", "апробация", "калейдоскоп", "шизофрения", "меланхолия",
	}
	wordMap = map[Difficulty]WordList{
		Easy:   easyWords,
		Medium: mediumWords,
		Hard:   hardWords,
	}
)

func (list WordList) RandomWord() string {
	index := rand.IntN(len(list))
	return list[index]
}

func GetWordList(difficulty Difficulty) (WordList, error) {
	wordList, ok := wordMap[difficulty]
	if !ok {
		return nil, errors.New("некорректная сложность игры")
	}
	return wordList, nil
}
