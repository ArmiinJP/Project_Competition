package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	n, aErr := strconv.Atoi(input)
	if aErr != nil {
		log.Fatalf("number is false")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	paper := scanner.Text()

	papers := strings.Split(paper, " ")
	var papersNum []int

	for _, v := range papers {
		vInt, _ := strconv.Atoi(v)
		papersNum = append(papersNum, vInt)
	}

	var hIndex = 0
	result := make(map[int]int)
	for i := n ; i > 0 ;i--{
		for _, v := range papersNum{
			if v >= i{
				hIndex++
			}
		}
		result[i] = hIndex
		if i < hIndex {
			break
		}
		hIndex = 0
	}

	max := 0
	for k, v := range(result){
		if k <= v{
			if k > max{
				max = k
			}
		}
	}

	//fmt.Println(result)
	fmt.Println(max)
}
