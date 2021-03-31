package main

import "fmt"

func Check(s string,ch rune) bool {
	i := 0

	arr := []rune(s)

	for i < len(s)-1 {
		if arr[i] == ch{
			break
		}
		i++
	}

	if arr[i] == ch {
		return true
	} else {
		return false
	}
}
func main() {
	s := "Marvellous Infosystem"
	ch := '0'
	bret := false
	
	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter Character You Want To Check")

	fmt.Scanf("%c",&ch)

	bret = Check(s,ch)
	if bret == true {
		fmt.Println("Character is There")
	} else {
		fmt.Println("Character is Not There")
	}
}