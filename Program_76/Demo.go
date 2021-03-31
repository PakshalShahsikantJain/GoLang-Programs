package main

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0
	j := 0

	ch := 'a'

	fmt.Print("Output :")
	for i = 0; i < Row;i++{
		fmt.Println()
		for j = 1;j <= Col;j++{
			if i % 2 == 0{
				fmt.Printf("%c\t",ch)
				ch++
			}else{
				fmt.Print(j,"\t")
			}
		}
		ch = 'a'
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

	DisplayPattern(No,No2)
}
