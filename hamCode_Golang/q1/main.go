package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputSlice := strings.Split(input, " ")

	var inputNum []int
	for _, v := range inputSlice {
		tmpNum, _ := strconv.Atoi(v)
		inputNum = append(inputNum, tmpNum)
	}

	fmt.Println(24 - ((inputNum[0] + inputNum[1]) - inputNum[2]))

}
