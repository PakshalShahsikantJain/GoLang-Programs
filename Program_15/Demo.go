package main

import "fmt"

func check(Ch byte) bool {
	if Ch == 'a' || Ch == 'e' || Ch == 'i' || Ch == 'o' || Ch == 'u' {
		return true
	} else {
		return false
	}
}
func main() {

	var Ch byte = '\x00'
	fmt.Println("Ganapati Bappa Morya...")
	fmt.Println("Enter Vowel To Check \t")
	fmt.Scanf("%c", &Ch)

	var bret bool = check(Ch)
	if bret == true {
		fmt.Println("Entered Character is Vowel")
	} else {
		fmt.Println("Entered Character is Not Vowel")
	}
}
