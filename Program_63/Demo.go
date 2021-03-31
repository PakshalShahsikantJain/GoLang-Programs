package main 

import "fmt"

func DisplayPattern(No int)  {
	i := 0
	Value := 1

	fmt.Println("Output :")
	for i = 1; i <= No ;i++{
		Value = i * 2
		fmt.Print(Value,"\t")
	}
}
func main(){
	No := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	DisplayPattern(No)
}