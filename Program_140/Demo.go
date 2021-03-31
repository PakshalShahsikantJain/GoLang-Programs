package main

import "fmt"

func FirstOccurance(s string,ch rune) int {
	i  := 0

	arr := []rune(s)

	for i = 0;i < len(s)-1;i++ {
		if arr[i] == ch {
			break
		}
	}

	if arr[i] == ch {
		return i
	} else {
		return -1
	}
}
func main() {
	s := "Marvellous Infosystem"
	ret := 0

	ch := '0'
	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Character To Find FirstOccurance")
	fmt.Scanf("%c",&ch)

	ret = FirstOccurance(s,ch)
	fmt.Println("FirstOccurance is At :",ret)
}