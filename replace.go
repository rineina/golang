package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string{

	new := strings.Replace(dna, "T", "U", -1)

	fmt.Println(dna)
	fmt.Println(new)
	return new
}

func main(){
	fmt.Println(DNAtoRNA("TTCAG"))
}

