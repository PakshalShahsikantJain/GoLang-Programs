package main

import "fmt"

func DisplayTime(ch rune) {
	if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
		if ch == 'A' || ch == 'a' {
			fmt.Println("Your Exam is At 7:00 AM")
		} else if ch == 'B' || ch == 'b' {
			fmt.Println("Your Exam is At 8:30 AM")
		} else if ch == 'C' || ch == 'c' {
			fmt.Println("Your Exam is At 9:30 AM")
		} else if ch == 'D' || ch == 'd' {
			fmt.Println("Your Exam is At 10:30 AM")
		} else {
			fmt.Println("Your Division is Invalid")
		}
	}
}


func main() {
	ch := '0'

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Your Division....")
	fmt.Scanf("%c",&ch)

	DisplayTime(ch)
}