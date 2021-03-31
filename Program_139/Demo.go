package main

import "fmt"

func Count(s string,ch rune) int {
	i := 0
	iCnt := 0

	arr := []rune(s)
	for i < len(s) {
		if arr[i] == ch {
			iCnt++
		}
		i++
	}

	return iCnt
}
func main() {
	s := "Marvellous Multi OS"
	ch := '0'

	ret := 0
	fmt.Println("Jay Ganesh........")

	fmt.Println("Enter Character To Count Frequency of")
	fmt.Scanf("%c",&ch)

	ret = Count(s,ch)
	fmt.Println("Frequency of Character is :",ret)
}