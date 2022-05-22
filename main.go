package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 1, 2, 3}
	// array := make([]int, 10)

	//fmt.Println(len(array))

	//array = append(array, 321)
	//array[2] = 12

	//sort.Ints(array)

	//fmt.Println(len(array))

	// for _, i:= range array {

	// 	fmt.Println(i)
	// }
	help := []int{}

	for i := 0; i < len(array); i++ {

		if array[i] == 2 {

			help = append(help, 2)
		}
	}

	fmt.Println(help)

	// if  condition-1 {
	// 	// code to be executed if condition-1 is true
	// } else if condition-2 {
	// 	// code to be executed if condition-2 is true
	// } else if condition-2 {
	// 	// code to be executed if condition-2 is true
	// } else {
	// 	// code to be executed if both condition1 and condition2 are false
	// }

	//	fmt.Println(array[1])
}
