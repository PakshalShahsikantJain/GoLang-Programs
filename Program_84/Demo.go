package main

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0
	j := 0

	fmt.Print("Output :")
	for i = 1; i <= Row;i++{
		fmt.Println()
		for j = 1; j <= Col;j++{
			if i == j || i > j {
				fmt.Print("*\t")
			} else {
				fmt.Print(" ")
			}
			//fmt.Print(i,j,"\t")
		}
	}
}
func main(){
	No := 0
	No2 := 0

	fmt.Println("Jay Ganesh...")
	
	fmt.Println("Enter Numebr of Rows")
	fmt.Scan(&No)

	fmt.Println("Enter Numebr of Columns")
	fmt.Scan(&No2)

	DisplayPattern(No,No2)
}