package main 

import "fmt"

func DisplayPattern(Row,Col int){
	i := 0
	j := 0

	fmt.Print("Output :")
	for i = Row ;i > 0;i--{
		fmt.Println()
		for j = 1; j <= Col;j++{
			if i == j || i == 1 || j == 1 || i == Row || j == Col{
				fmt.Print("*\t")
			} else if i > j {
				fmt.Print("#\t")
			} else {
				fmt.Print("$\t")
			}
			//fmt.Print(i,j,"\t")
		}
	}
}
func main(){
	No := 0
	No2 := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Number of Rows")
	fmt.Scan(&No)

	fmt.Println("Enter Number of Columns")
	fmt.Scan(&No2)

	DisplayPattern(No,No2)
}