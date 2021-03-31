package main

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0
	j := 0

	fmt.Print("Output :")
	for i = 0; i < Row;i++{
		fmt.Println()
		for j = 0; j < Col;j++{
			if j % 2 == 0{
				fmt.Print("*\t")
			} else {
				fmt.Print("#\t")
			}
		}
	} 
}

func main(){
	No := 0
	No2 := 0

	fmt.Println("Jay Ganesh.....")

	fmt.Println("Enter Number of Rows")
	fmt.Scan(&No)

	fmt.Println("Enter Number of Columns")
	fmt.Scan(&No2)

	DisplayPattern(No,No2)
}