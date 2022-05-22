package main

import "fmt"

func main(){
	defer finish()
	defer fmt.Println("Program has been started")
	fmt.Println("Program working")
}

func finish(){
	fmt.Println("Program has been finished")
}
