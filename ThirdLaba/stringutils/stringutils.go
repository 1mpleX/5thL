package stringutils

import "fmt"

func ReverseString() {
	fmt.Print("\n\ntask_3\nEnter text: ")

	var a string
	fmt.Scan(&a)

	arr := []rune(a)

	fmt.Print("reversed text - ")
	for b := len(arr) - 1; b >= 0; b-- {
		fmt.Print(string(arr[b]))
	}
}
