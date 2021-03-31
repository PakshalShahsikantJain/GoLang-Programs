package main

import "fmt"

func Check(No,No2 int) bool {
	Output := 0
	Mask := 0X00000001

	if No < 0 {
		No = -No
	}

	if (No2 < 1) || (No2 > 32) {
		return false
	} 

	Mask = Mask << No2 
	Output = No & Mask

	if Output == Mask {
		return true 
	} else {
		return false
	}
}
func main() {
	No := 0
	No2 := 0
	bret := false 

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter Position To Check")
	fmt.Scan(&No2)

	bret = Check(No,No2)

	if bret == true {
		fmt.Println("Bit is On")
	} else {
		fmt.Println("Bit is Off")
	}
}