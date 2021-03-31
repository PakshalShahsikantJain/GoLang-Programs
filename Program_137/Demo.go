package main

import "fmt"

func Count(s string) int {
	i := 0
	iCnt := 0

	arr := []rune(s)
	for i < len(s) {
		if arr[i] == ' '{
			iCnt++
		}
		i++
	}

	return iCnt
}
func main(){
	//s := "Marvellous Infosystem by Piyush Manohar Khairnar"
	s2 := "Pakshal Karande"

	ret := 0

	fmt.Println("Jay Ganesh......")

	ret = Count(s2)
	fmt.Println("No of White Spaces are :",ret)
}