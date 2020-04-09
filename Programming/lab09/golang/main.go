// RUS С клавиатуры вводится текст, состоящий из строк (не более 20).
// Длина каждой строки - не более 128. Слова разделены одним или
// более пробелами, в каждой строке не менее 2 слов.
// Задача: удалить из каждой строки слова с чётными номерами.

// ENG Text, consisting of strings (less than 20), is being printed from keyboard.
// Every stroke has a length less than 128. Words divided by one or more
// spaces, number of words - more than 2.
// Objective is delete words with even numbers from every stroke.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, i, spaceC, l int
	str             [20]string
	words           [65]string
)

func main() {
	input()
	for i = 0; i < n; i++ {
		divide()
		deleter()
		for j := 0; j <= spaceC; j++ {
			if j%2 == 0 {
				fmt.Print(words[j], " ")
			}
		}
		fmt.Print("\n")

	}
}

func input() {
	fmt.Print("Введите количество строк: ")
	fmt.Scan(&n)
	if n > 20 {
		fmt.Print("Допустимый лимит - 20 строк!")
		os.Exit(1)
	} else {
		for i := 0; i < n; i++ {
			str[i], _ = bufio.NewReader(os.Stdin).ReadString('\n')
			if len(str[i]) > 128 {
				fmt.Println("Допустимый лимит - 128 символов!")
				os.Exit(1)
			}
		}
	}
}

func divide() {
	spaceC = -1
	l = len(str[i])
	for j := 0; j < l; j++ {
		if string(str[i][j]) == " " {
			spaceC++
			words[spaceC] = str[i][0:j]
			str[i] = str[i][j+1:]
			l = len(str[i])
			j = 0
		}
		if string(str[i][j]) == "\n" {
			spaceC++
			words[spaceC] = str[i][0:j]
			str[i] = ""
		}
	}

}

func deleter() {
	for j := 0; j <= spaceC; j++ {
		if words[j] == " " {
			for w := j; w < spaceC; w++ {
				words[w] = words[w+1]
			}
		}
	}
}
