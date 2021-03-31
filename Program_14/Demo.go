package main

import "fmt"

func DisplayConvert(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch = ch + 32
	} else if ch >= 'a' && ch <= 'z' {
		ch = ch - 32
	}

	return ch
}

func main() {
	var Ch byte = '\x00'

	fmt.Println("Enter Character To Convert")
	fmt.Scanf("%c", &Ch)

	var ret byte = DisplayConvert(Ch)
	fmt.Printf("Converted Character is :  %c \n", ret)
}
