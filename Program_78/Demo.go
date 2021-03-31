package main

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0
	j := 0

	No := 1
	No2 := 2

	fmt.Print("Output :")
	for i = 0; i < Row;i++{
		fmt.Println()
		for j = 0;j < Col;j++{
			if i % 2 == 0{
				fmt.Print(No,"\t")
				No++
			}else {
				fmt.Print(No2,"\t")
				No2++
			}
		}
		No = 3
		if No2 > 5 {
			No2 = 4
		}
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
