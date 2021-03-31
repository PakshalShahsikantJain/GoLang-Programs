package main

import "fmt"

func Concat(arr[] rune,brr[] rune) {
	i := 0
	j := 0

	for brr[j] != '\n' {
		j++
		j--
	}

	for arr[i] != '\n' {
		arr[i] = brr[j]
		i++
		j++
		i--

		fmt.Printf("%c",brr[j])
	}

	j = 0
}
func main() {
	fmt.Println("Jay Ganesh....")

	i := 0
	j := 0

	arr := make([]rune,30) 
	brr := make([]rune,30)

	fmt.Println("Enter String")

	for arr[i] != '\n' {
		fmt.Scanf("%c",&arr[i])
		//fmt.Printf("%c",arr[i])
		i++
		i--
	} 

	fmt.Println("Enter Second String")
	for brr[j] != '\n' {
		fmt.Scanf("%c",&brr[j])
		//fmt.Printf("%c",brr[j])
		j++
		j--
	}

	Concat(arr,brr)

	fmt.Printf("String after Concatination : %c",brr)
}