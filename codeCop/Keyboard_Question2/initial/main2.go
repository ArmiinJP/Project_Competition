package keyboard

//package main

import (
	// "bufio"
	// "fmt"
	// "log"
	// "os"
	// "strconv"
	"strings"
	"unicode"
)

type Keyboard struct {
	quality	int
	click map[rune]int
	delete []rune
	result string
}

func NewKeyboard(dure int) *Keyboard {
	return &Keyboard{
		quality: dure,
		delete: []rune{},
		click: map[rune]int{},
		result: "",
	}
}

func (keyboard *Keyboard) Enter(inp string) string {
	
	var testShift int

	newInp := ""
	for newInp == inp{
		newInp = strings.Replace(inp, "\n", "", 3)
	}
	for _, key := range inp{

		check := true
		if unicode.IsUpper(key){
			testShift += 1
			if testShift > keyboard.quality {
				key = unicode.ToLower(key)
			}
		}


		// fmt.Println(unicode.IsUpper(key), key, string(key))

		
		keyboard.click[unicode.ToLower(key)] += 1
		if keyboard.click[unicode.ToLower(key)] == keyboard.quality{
			keyboard.delete = append(keyboard.delete, unicode.ToLower(key))
			keyboard.result += string(key)
			
			continue
		}

		//fmt.Println(key)
		for _, v  := range keyboard.delete{
			if unicode.ToLower(key) == v{
				check = false
				break
			}
		}
		
		if check {
			keyboard.result += string(key)
		}

	}

	return keyboard.result
}

// func main(){
// 	var input string
// 	fmt.Scanln(&input)
	
// 	n, aErr := strconv.Atoi(input)
// 	if aErr != nil {
// 		log.Fatalf("number is false")
// 	}

// 	keyboard := NewKeyboard(n)

// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()

// 	fmt.Println(keyboard.Enter(scanner.Text()))
	
// }