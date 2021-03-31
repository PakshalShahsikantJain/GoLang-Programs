package main

import "fmt"

func Count(s string) int {
	i := 0
	iCnt := 0

	arr := []rune(s)

	for i < len(s) {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			iCnt++
		}
		i++
	}

	return iCnt
}
func main() {
	s := "Marvellous Multi OS"
	ret := 0

	fmt.Println("Jay Ganesh.......")

	ret = Count(s)
	fmt.Println("Number of Capital Letters in String are :",ret)
}