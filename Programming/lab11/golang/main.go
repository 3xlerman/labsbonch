package main

import (
	"fmt"
	"io/ioutil"
)

var (
	enter int
	str   [50]string
	words [50][10]string
	data  []byte
)

func main() {
	input()
	divideToStrings()
	divideToWords()
	output()
}

func input() {
	data, _ = ioutil.ReadFile("numbers.txt")
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
	for i := 0; i < enter; i++ {
		wordsC := 0
		l := len(str[enter])
		for j := 0; j < l; j++ {
			if string(str[i][j]) == " " {
				words[i][wordsC] = str[i][:j]
				str[i] = str[i][j+1:]
				wordsC++
				fmt.Println(words[i][wordsC])
				j = 0
				l = len(str[enter])
			}

		}
	}
}
func output() {
	for i := 0; i < enter; i++ {
		fmt.Println(str[i])
	}
}
