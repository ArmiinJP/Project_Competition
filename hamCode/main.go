package main

import (
	"fmt"
	"log"
	"strconv"
)

func  main(){
	var input string
	fmt.Scanln(&input)

	inputNum, aErr := strconv.Atoi(input)
	if aErr != nil{
		log.Fatalf("number is false")
	}

	for i := 0 ; i < inputNum ; i++{
		fmt.Printf("#")
	}
	fmt.Println()
}