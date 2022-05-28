package main

import (
	"fmt"
	"strings"
)

func main(){
   word := strings.Split("cat", "")
  
   var reverse []string

   for i := len(word) - 1; i >= 0; i-- {
      reverse = append(reverse, word[i])
   //fmt.Println(word[i])
   
}
fmt.Println(strings.Join(reverse, ""))



   // Вывод: [5 4 3 2 1]
   fmt.Println(reverse)
}