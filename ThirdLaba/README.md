Создать пакет mathutils с функцией для вычисления факториала числа.
Использовать созданный пакет для вычисления факториала введенного пользователем числа.
Создать пакет stringutils с функцией для переворота строки и использовать его в основной программе.
Написать программу, которая создает массив из 5 целых чисел, заполняет его значениями и выводит их на экран.
Создать срез из массива и выполнить операции добавления и удаления элементов.
Написать программу, которая создает срез из строк и находит самую длинную строку.

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

![alt text](image.png)

запускаемый файл - main.go
или же в терминале go run "путь до файла с расширением"