package main

import "fmt"

func main() {

	str1 := " "
	str2 := " "
	fmt.Println("Jay Ganesh....")

	fmt.Println("Enter Your name")
	fmt.Scanf("%s %s", &str1, &str2)

	fmt.Printf("Your name is %s \n", str1+" "+str2)
}
