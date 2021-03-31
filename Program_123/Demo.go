package main

import "fmt"

func Display() {
	i := 0

	fmt.Println()
	for i = 0;i <= 255;i++ {
		fmt.Printf(" %d  %x  %o  %c",i ,i ,i ,i)
		fmt.Println()
	}
}
func main() {
	fmt.Println("Jay Ganesh.........")
	Display()
}