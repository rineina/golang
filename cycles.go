package main

import "fmt"

func main(){
	i := 5
	for i <= 10{
		fmt.Println(i)
		i = i + 1
	}

	for i:=0; i <=5; i++ {
		fmt.Println(i)
	}

	text := []string{"be", "mouse", "hello"} 
	
	for _, i := range text {
		
		if i == "mouse"{
			continue}
			fmt.Println(i)
	}

}






