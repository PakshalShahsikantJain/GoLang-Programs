package main 

import "fmt"

func main() {
	i := 0

	fmt.Println("Jay Ganesh.....")
	arr := make([]rune,30) 
	brr := make([]rune,30)

	fmt.Println("Enter String")
	for arr[i] != '\n' {
		fmt.Scanf("%c",&arr[i])
		brr[i] = arr[i]

		if brr[i] >= 'A' && brr[i] <= 'Z' {
			fmt.Printf("%c",brr[i])
		}
		i++
		i--
	}
}