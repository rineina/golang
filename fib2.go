package main

import "fmt"

func fibbonachi(n int) int {
	if n < 2 {
		return n
	}
	return fibbonachi(n - 2) + fibbonachi(n - 1)
}

func main(){
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibbonachi(i))
	}
	fmt.Printf("\n")
	var h int 
	for h = 0; h < 9; h++ {
		fmt.Printf("%d\t", fibbonachi(h))
	}
}