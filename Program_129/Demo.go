package main

import "fmt"

func count(s string) int {
	i := 0
	iCnt := 0

	for i < len(s) {
		arr := []rune(s)

		if arr[i] >= 'a' && arr[i] <= 'z' {
			iCnt++
		}
		i++
	}

	return iCnt
}
func main() {
	s := ""
	ret := 0

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter String")
	fmt.Scan(&s)

	ret = count(s)

	fmt.Println("Number of small letters are :",ret)
}