package main

import "fmt"

func main() {
	names := [3]string{
		"Marvellous",
		"Infosystem",
		"Logic",
	}
	names2 := [3]string{"Pakshal","Infosystem","Logic"} 

	fmt.Println("Jay Ganesh........")

	fmt.Println(names)
	fmt.Println(names2)

	if names[0] == names2[0] {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}