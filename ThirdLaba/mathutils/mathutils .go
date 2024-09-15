package mathutils

import "fmt"

//task_1
func Fact() {
	fmt.Print("\ntask_1\nEnter number: ")
	var a int64
	sum := int64(1)

	fmt.Scan(&a)

	if a == 0 {
		fmt.Print(1)
	} else if a > 0 {
		for b := int64(1); b <= a; b++ {
			sum *= b
		}
		fmt.Print("fact ", a, " - ", sum)
	} else {
		fmt.Print("Enter the positive number")
	}

}
