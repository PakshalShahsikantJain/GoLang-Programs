package main

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0

	fmt.Print("Output :")
	for i = Row; i > 0;i--{
		fmt.Println()
		for j := Col; j > 0;j--{
			fmt.Print(j,"\t")
		}
	}
}
func main(){
	No := 0
	No2 := 0
	fmt.Println("Jay Ganesh.......")

	fmt.Println("Enter Number of Rows")
	fmt.Scan(&No)

	fmt.Println("Enter Number of Columns")
	fmt.Scan(&No2)

	DisplayPattern(No,No2)
}