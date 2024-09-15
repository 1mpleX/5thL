package main

import (
	"GOLABI/ThirdLaba/mathutils"
	"GOLABI/ThirdLaba/stringutils"
	"fmt"
)

func main() {
	mathutils.Fact()
	stringutils.ReverseString()
	AddArr()
	SliceStringArr()
}

// task_4
func AddArr() {
	arr := []int{1, 2, 3, 4, 5}
	for a := 10; a < 15; a++ {
		arr = append(arr, a)
	}
	fmt.Print("\n\ntask_4\n", arr)
	SliceArr(arr)
	fmt.Print(SliceArr(arr))
}

// task_5
func SliceArr(a []int) []int {
	slice := a[2:6]
	fmt.Print("\n\ntask_5\n", "sliced array: ", slice)
	slice = append(slice, 20)
	fmt.Print("\nadded 20 ", slice)
	fmt.Print("\ndeleted 1st elem ")
	return slice[1:]
}

// task_6
func SliceStringArr() {
	print("\n\ntask_6\n")
	var first, second, third string

	fmt.Print("Enter 1st text: ")
	fmt.Scan(&first)
	fmt.Print("Enter 2nd text: ")
	fmt.Scan(&second)
	fmt.Print("Enter 3rd text: ")
	fmt.Scan(&third)

	arr := []string{first, second, third}
	fmt.Print("\nyour array: ", arr)

	arr[1] = "$uicideboy$"
	fmt.Print("\nupdated array: ", arr)

	myMax := int(0)
	for i := int(0); i < len(arr); i++ {
		if len(arr[i]) > len(arr[myMax]) {
			myMax = i
		}
	}
	print("\nlongest line - ", arr[myMax])
}
