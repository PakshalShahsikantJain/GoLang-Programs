package main

import "fmt"

func TableRev(No int) {
	i := 0
	Table := 0

	if No < 0 {
		No = -No
	}

	fmt.Print("Output : ")
	for i = 10; i > 0; i-- {
		Table = No * i

		fmt.Print(Table, "\t")
	}
}
func main() {
	No := 0
	fmt.Println("Jay Ganesh....")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	TableRev(No)
}
