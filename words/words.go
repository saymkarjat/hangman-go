package words

import "errors"

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

type WordList []string

var (
	EasyWords = WordList{
		"кот", "дом", "мяч", "сон", "лес",
	}
	MediumWords = WordList{
		"корабль", "молоток", "зеркало", "черепаха", "крокодил",
	}
	HardWords = WordList{
		"дихотомия", "апробация", "калейдоскоп", "шизофрения", "меланхолия",
	}
	wordMap = map[Difficulty]WordList{
		Easy:   EasyWords,
		Medium: MediumWords,
		Hard:   HardWords,
	}
)

func GetWordList(difficulty Difficulty) (WordList, error) {
	wordList, ok := wordMap[difficulty]
	if !ok {
		return nil, errors.New("Ne nayden level")
	}
	return wordList, nil
}
