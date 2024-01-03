package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n###### Start sutom helper ######")

	searchWord := questionaire()
	listWords := loadListOfWords("mots.txt")

	filtreByLength := filterByWordLength(searchWord.length, listWords)
	filtreList := filtreByLetterAndPosition(searchWord.knowLetterWithPosition, 0, filtreByLength)

	if len(searchWord.letterNotInThisWord) != 0 {
		filtreList = filterWordWithMissingLetters(searchWord.letterNotInThisWord, 0, filtreList)
	}

	for _, wordInList := range filtreList {
		fmt.Println("\t- ", string(wordInList))
	}
}
