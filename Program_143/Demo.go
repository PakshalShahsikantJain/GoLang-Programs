package main

import "fmt"

func main() {
	i := 0
	No := 10

	fmt.Println("Jay Ganesh........")

	arr := make([]rune,30) 
	brr := make([]rune,30)

	fmt.Println("Enter String")

	for arr[i] != '\n' && No != 0{
		fmt.Scanf("%c",&arr[i])
		brr[i] = arr[i]
		i++
		i--
		No--
		fmt.Printf("%c",brr[i])
	}
}
