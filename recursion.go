package main

import "fmt"

func factorial(n uint) uint{

	if n == 0{
		return 1
	}
	return n * factorial(n - 1)
}
func main(){
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
	fmt.Println(factorial(0))
}