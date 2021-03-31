package main

import "fmt"

func DispalyPattern(Row, Col int) {
	i := 0
	j := 0

	ch := 'A'

	fmt.Print("Output :")
	for i = 0; i < Row; i++ {
		fmt.Println()
		for j = 0; j < Col; j++ {
			fmt.Printf("%c", ch)
			ch++
		}
		ch = 'A'
	}
}
func main() {
	No := 0
	No2 := 0

	fmt.Println("Jay Ganesh.......")

	fmt.Println("Enter Number of Rows")
	fmt.Scan(&No)

	fmt.Println("Enter Number of Columns")
	fmt.Scan(&No2)

	DispalyPattern(No, No2)
}
