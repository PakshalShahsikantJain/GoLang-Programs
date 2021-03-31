package main

import "fmt"

func Check(s string) bool {
	i := 0

	arr := []rune(s)

	for i < len(s) - 1 {
		if arr[i] == 'a' || arr[i] == 'e'|| arr[i] == 'i' || arr[i] == 'o' || arr[i] == 'u' {
			break
		} 
		i++ 
	}

	if arr[i] == 'a' || arr[i] == 'e' || arr[i] == 'i' || arr[i] == 'o' || arr[i] == 'u' {
		return true
	} else {
		return false
	}
 }
func main() {
	s := ""
	bret := false
	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter String")
	fmt.Scan(&s)

	fmt.Println(s)

	bret = Check(s)

	if bret == true {
		fmt.Println("String Contains Vowels")
	} else {
		fmt.Println("String Does Not Contains Vowels")
	}
}