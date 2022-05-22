package main

import (
	"fmt"
)

func main(){

	const mouse string = "mouse"
	textHello := "Hello"
	textBe := "Be"
	
	i :=10


	fmt.Println(fmt.Sprintf("%s %d %s", textHello, i, textBe), mouse)

	var input float64
	fmt.Scanf("%f", &input)
	output :=input + 1
	fmt.Println(output)

}






