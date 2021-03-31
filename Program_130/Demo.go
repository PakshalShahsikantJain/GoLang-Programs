package main

import "fmt"

func Diff(s string) int {
	i := 0
	iCnt := 0
	iCnt2 := 0

	Difference := 0

	for i < len(s) {
		arr := []rune(s)

		if arr[i] >= 'A' && arr[i] <= 'Z' {
			iCnt++ 
		} else {
			iCnt2++
		}
		i++
	}

	Difference = iCnt2 - iCnt

	return Difference
}
func main() {
	s := ""
	ret := 0

	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter String")
	fmt.Scan(&s)

	ret = Diff(s)
	fmt.Println("Difference is  :",ret)
}