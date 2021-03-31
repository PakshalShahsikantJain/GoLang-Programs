package main

import "fmt"

func WhiteSpace(s string) {
	
	arr := []rune(s)
	brr := make([]rune,30)

	fmt.Println("Output :")
	for i := 0;i < len(s);i++ {
		if arr[i] == ' '{
			brr[i-1] = arr[i-1]
		} else {
			brr[i] = arr[i]
		}

		fmt.Printf("%c",brr[i])
	}
}
func main() {
	s := "Marve llous Mu lti O S"

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Input :")
	fmt.Println(s)

	WhiteSpace(s)
}