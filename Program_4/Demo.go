package main

import "fmt"

func Divide(No1, No2 float64) float64 {

	if No2 <= 0 {
		return -1
	}
	ans := No1 / No2

	return ans

}

func main() {

	var No float64
	var No2 float64

	fmt.Println("Enter First Number")
	fmt.Scan(&No)

	fmt.Println("Enter Second Number")
	fmt.Scan(&No2)

	ret := Divide(No, No2)

	fmt.Println(ret)
}
