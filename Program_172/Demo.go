package main

import "fmt"

func Pattern(s string) {
	i := 0
	j := 0

	arr := []rune(s)

	for i = 0;i < len(arr);i++ {
		fmt.Println()
		for j = 0;j < len(arr);j++ {
			if i == j || i > j {
				fmt.Printf("%c\t",arr[j])
			}
		}
	}
}
func main() {
	s := "PPA"

	fmt.Println("Jay Ganesh.........")

	Pattern(s)
}