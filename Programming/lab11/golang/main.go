package main

import (
	"fmt"
	"io/ioutil"
)

var (
	enter int
	str   [50]string
	words [10]string
	data  []byte
)

func main() {
	data, _ = ioutil.ReadFile("numbers.txt")
	divideToStrings()
	//divideToWords()
	output()
}

func output() {
	for i := 0; i < enter; i++ {
		fmt.Println(str[i])
	}
}

func divideToStrings() {
	enter = 0
	for i := 0; i < len(data); i++ {
		if data[i] == byte('\n') {
			str[enter] = string(data[:i])
			data = data[i+1:]
			enter++
		}
	}
	str[enter] = string(data)
	enter++
}

func divideToWords() {
	spaceC := -1
	for i := 0; i < enter; i++ {
		l := len(str[i])
		for j := 0; j < l; j++ {
			if string(str[i][j]) == " " { // dividing a stroke in words
				spaceC++
				words[spaceC] = str[i][0:j]
				str[i] = str[i][j+1:]
				l = len(str[i])
				j = 0
			}
			if string(str[i][j]) == "\n" { // adding last word in array of words
				spaceC++
				words[spaceC] = str[i][0:j]
				str[i] = ""
			}
		}
	}

}
