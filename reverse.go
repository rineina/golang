package main

import "fmt"

func Solution(word string) string {
	array := []rune{word}
	for i, j := 0; len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return string(word)
}


func main(){
	fmt.Println()
}