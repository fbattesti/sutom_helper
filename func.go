package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func checkErr(err error, message string) {
	if err != nil {
		fmt.Println("ERROR: "+message, err)
		os.Exit(1)
	}
}

func remove_last_caract(mystring string) string {
	size_string := len(mystring)
	if size_string > 0 {
		mystring = mystring[:size_string-1]
		return mystring
	} else {
		return mystring
	}
}

func questionaire() (searchWord searchWord) {

	fmt.Println("\nVeuillez renseigner les lettres connus avec leur position sour la forme ")
	fmt.Println("Exemple: s-to  pour le mot sutom avec des - pour les emplacements inconnu et ne rajouter pas des - a la fin")
	fmt.Print("> ")

	readInput := bufio.NewReader(os.Stdin)
	knowLetterWithPosition, err := readInput.ReadString('\n')
	checkErr(err, "Input knowLetterWithPosition")
	searchWord.knowLetterWithPosition = remove_last_caract(knowLetterWithPosition)

	fmt.Println("Veuillez renseigner le nombre de caractere du mot a trouver")
	fmt.Print("> ")
	readInput = bufio.NewReader(os.Stdin)
	lengthWord, err := readInput.ReadString('\n')
	checkErr(err, "Input lengthWord")
	// Convert string to int + remove the space after the value before convert
	searchWord.length, _ = strconv.Atoi(remove_last_caract(string(lengthWord)))

	fmt.Println("Veuillez renseigner les lettres connues ( sans importance d ordre ex  : abcz ) ")
	fmt.Print("> ")
	readInput = bufio.NewReader(os.Stdin)
	searchWord.knowLetter, err = readInput.ReadString('\n')
	checkErr(err, "Input knowLetter")

	searchWord.numberKnowLetter = len(string(searchWord.knowLetter)) - 1

	fmt.Println("Veuillez renseigner les lettres absentes ( sans importance d ordre ex  : abcz ) ")
	fmt.Print("> ")
	readInput = bufio.NewReader(os.Stdin)
	lettersNotInThisWord, err := readInput.ReadString('\n')
	checkErr(err, "Input letterNotInThisWord")

	// remove letter already set in know letter in letterNotInThisWord string
	for _, letter := range lettersNotInThisWord {
		if strings.Contains(searchWord.knowLetter, string(letter)) || strings.Contains(searchWord.knowLetterWithPosition, string(letter)) {
			//Do nothing
		} else {
			searchWord.letterNotInThisWord = searchWord.letterNotInThisWord + string(letter)
		}
	}

	return searchWord
}

func loadListOfWords(FilePath string) (listWords []string) {
	// Read the file
	dict, err := os.ReadFile(FilePath)
	checkErr(err, "ReadFile : "+FilePath)
	// make []string, one string by line
	listWords = strings.Split(string(dict), "\n")
	return listWords
}

func filterByWordLength(searchLength int, listWords []string) (filtreList []string) {
	for _, word := range listWords {
		if utf8.RuneCountInString(word) == searchLength {
			filtreList = append(filtreList, word)
		}
	}
	return filtreList
}

func filterWordWithThisLetterInThisPosition(positionLetter int, letter string, listWords []string) (filtreList []string) {
	for _, word := range listWords {
		if word[positionLetter] == letter[0] {
			filtreList = append(filtreList, word)
		}
	}
	return filtreList
}

func filtreRemoveWordWithThisLetter(letter string, listWords []string) (filtreList []string) {
	var letterIsMissing bool
	for _, word := range listWords {
		letterIsMissing = true
		if strings.Contains(word, letter) {
			letterIsMissing = false
		}
		if letterIsMissing {
			filtreList = append(filtreList, word)
		}
	}
	return filtreList
}

// Recursive func
func filterWordWithMissingLetters(letters string, iteration int, listWords []string) (filtreList []string) {
	// if iteration = nomber of letter stop recurisv func
	if utf8.RuneCountInString(letters) <= iteration {
		filtreList = listWords
		return filtreList
	} else {
		filtreList = filtreRemoveWordWithThisLetter(string(letters[iteration]), listWords)
		iteration++
		return filterWordWithMissingLetters(letters, iteration, filtreList)
	}
}

// Recursive func
func filtreByLetterAndPosition(knowLetterWithPosition string, nombreIteration int, listWords []string) []string {

	// Stop condition: If knowLetterWithPosition is empty, return the current list
	if utf8.RuneCountInString(knowLetterWithPosition) == 0 {
		return listWords
	}

	var filtre []string
	for i_positionInWord, char := range knowLetterWithPosition {
		if char != '-' {
			filtre = filterWordWithThisLetterInThisPosition(i_positionInWord+nombreIteration, string(char), listWords)
			break // Casser la boucle après la première substitution

		}
	}

	//recursive call with filtered list and string knowLetterWith Position updated
	nombreIteration++
	return filtreByLetterAndPosition(knowLetterWithPosition[1:], nombreIteration, filtre)
}
