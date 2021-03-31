package main

import "fmt"

func Mult(i, j int) int {

	ans := i * j
	return ans
}
func main() {

	fmt.Println("Enter First Number")
	var No int
	fmt.Scan(&No)

	fmt.Println("Enter Second Number")
	var No2 int
	fmt.Scan(&No2)

	ret := Mult(No, No2)

	fmt.Println(ret)
}
