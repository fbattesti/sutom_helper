
// Importing fmt
package main

// bufio pour le read
import (
	"fmt"
	"os"
	"bufio" 
	"strings"
	"strconv"
	
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func remove_last_caract (mystring string ) string {
	size_string := len(mystring)
	if size_string > 0 {
		mystring = mystring[:size_string-1]
		return mystring
	}else{
		fmt.Println("aucun espace a effacer")
		return mystring
	}
}


// Calling main
func main() {
  
	// #####################  Ask info for the script ##########################

	fmt.Println("Start sutom helper ")
	fmt.Println("Veuillez renseigner la premiere lettre du mot a trouver")
	fmt.Print("> ")
	
	reader := bufio.NewReader(os.Stdin)
	sutom, _ := reader.ReadString('\n')

	fmt.Println("Veuillez renseigner le nombre de caractere du mot a trouver")
	fmt.Print("> ")
	reader2 := bufio.NewReader(os.Stdin)
	len_sutom, _ := reader2.ReadString('\n')
	// Convert string to int + remove the space after the value before convert
	int_len_sutom, _ := strconv.Atoi(remove_last_caract(string(len_sutom)))

	fmt.Println("Veuillez renseigner les lettres connues ( sans importance d ordre ex  : abcz ) ")
	fmt.Print("> ")
	reader3 := bufio.NewReader(os.Stdin)
	letter_know, _ := reader3.ReadString('\n')
	fmt.Println("je sais qu il y a les lettres : ",string(letter_know))
	fmt.Println("je connais donc  : ",len(string(letter_know)),"lettre")
	number_know := len(string(letter_know))-1


	fmt.Println("Veuillez renseigner les lettres absentes ( sans importance d ordre ex  : abcz ) ")
	fmt.Print("> ")
	reader4 := bufio.NewReader(os.Stdin)
	letter_missing, _ := reader4.ReadString('\n')

	// Load dict
	dict, err := os.ReadFile("liste_francais.txt")
	check(err)
	list_word := strings.Split(string(dict),"\n")
	
 
	// For each word in dict
	for i, _ := range list_word  {
		// if start with the same letter 
		if list_word[i][0] == sutom[0] { 	
			// if there is same lenght
			if (len(string(list_word[i]))-1) == int_len_sutom {
				calcul_know_letter := 0
				for y, _ := range letter_know {
					if strings.Contains(list_word[i], string(letter_know[y]) ) {
						calcul_know_letter++
					}
				}
				// si les lettre connu sont dans le mot
				if calcul_know_letter == number_know {
					calcul_missing_letter := 0
					for yy, _ := range letter_missing {
						if strings.Contains(list_word[i], string(letter_missing[yy]) ) {
							calcul_missing_letter++
						}
					}
					if calcul_missing_letter == 0 {
						fmt.Println("C'est possiblement ",string(list_word[i]))
					}
				} // END if calcul_know_letter == number_know
			} // END same lenght
		} // END same first letter
	} // END for each word in dict	
} // END 