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

	inputNum, aErr := strconv.Atoi(input)
	if aErr != nil {
		log.Fatalf("number is false")
	}


	scanner := bufio.NewScanner(os.Stdin)
	var baze = []string{}
	for i := 0 ; i < inputNum ; i++{
		scanner.Scan()
		baze = append(baze, scanner.Text())
	}

	var elmIn, elmOut []string
	for i, _ := range(baze){
		elmIn = append(elmIn, strings.Split(baze[i],",")[0])
		elmOut = append(elmOut, strings.ReplaceAll(strings.Split(baze[i],",")[1], " ", ""))
	}

	var elmInNum, elmOutNum []int
	for _, v := range(elmIn){
		tmpString := strings.ReplaceAll(v, "[", "")
		tmpString = strings.ReplaceAll(tmpString, "(", "")
		a, _ := strconv.Atoi(tmpString)
		elmInNum = append(elmInNum, a)
	}	

	for _, v := range(elmOut){
		tmpString := strings.ReplaceAll(v, "]", "")
		tmpString = strings.ReplaceAll(tmpString, ")", "")
		a, _ := strconv.Atoi(tmpString)
		elmOutNum = append(elmOutNum, a)
	}	

	var elmInNumNew, elmOutNumNew []int

	tmp := elmInNum[0]
	tmpi := -1
	for i, v := range(elmInNum){
		
		if v < tmp {
			tmp = v
			tmpi = i 
		}
	}

	fmt.Println(elmIn)
	fmt.Println(elmInNum)
	fmt.Println(elmOutNum)
	fmt.Println(elmOut)



}


func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}