package main

import "fmt"

func main(){
	//var array [3]string
	//array := []int{1, 2, 3}
	
	array := make([]int, 10)

	fmt.Println(len(array))
	
	array = append(array, 321)
	
	for _, i := range array{
		fmt.Println(i)
	}

	fmt.Println(array)
}
