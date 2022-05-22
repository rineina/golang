package main

import "fmt"

func main(){
	keys := []string{}

	carta := map[string] []int{
		"1": {1, 2},
		"2": {2, 10},
		"cat": {666, 555},
	}


	fmt.Println(carta["cat"])


	for i, key := range carta {
		//keys = append(keys, key)
		fmt.Println(i, key)
	}
	fmt.Println(keys)
}
