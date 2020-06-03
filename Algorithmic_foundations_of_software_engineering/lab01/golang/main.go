package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	// L = lines
	L = 100

	// A = assignments
	A = 100

	// P = powers
	P = 100
)

var (
	// number of lines and assignments
	lineN, perN int

	// maximum power of every assign in every stroke
	maxP [L][A]int

	// F factors array
	F [L][A][P]float64

	// D derivative factors array
	D [L][A][P]float64

	// array of values
	values [A]float64

	// array of minimum values
	valuesMin [A]float64

	// array of gradient values
	grad [A]float64

	// epsilon
	epsilon float64
)

// detecting error type
func errCheck(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "expected integer") {
			fmt.Println("Неверный тип данных")
			return true
		}
		if strings.Contains(err.Error(), "value out of range") {
			fmt.Println("Выход за границы диапазона")
			return true
		}
		if strings.Contains(err.Error(), "invalid syntax") {
			fmt.Println("Неверный тип данных")
			return true
		}
	}
	return false
}

// const numbers input
func inputNumbers() {
	var temp float64

	fmt.Print("Количество строк в системе: ")
	_, err := fmt.Scan(&temp)
	if errCheck(err) {
		fmt.Println("Ошибка в вводе количества строк")
		os.Exit(1)
	}
	lineN = int(temp)

	fmt.Print("Количество переменных в системе: ")
	_, err = fmt.Scan(&temp)
	if errCheck(err) {
		fmt.Println("Ошибка в вводе количества переменных")
		os.Exit(1)
	}
	perN = int(temp)

	if lineN < perN {
		fmt.Println("Система неразрешима")
		os.Exit(1)
	}

}

// system factors input
func inputFactors() {
	var temp float64
	for i := 0; i < lineN; i++ {
		fmt.Print("### Строка ", i+1, ": ###\n")
		for j := 0; j < perN; j++ {

			fmt.Print("Переменная x[", j, "]:\n")
			fmt.Print("Высшая степень x[", j, "]: ")

			_, err := fmt.Scan(&temp)
			if errCheck(err) {
				fmt.Println("Ошибка в вводе высшей степени переменной x[", j, "]")
				os.Exit(1)
			}
			maxP[i][j] = int(temp)

			for k := maxP[i][j]; k >= 0; k-- {

				fmt.Print("Введите коэффициент при степени ", k, ": ")

				_, err = fmt.Scan(&F[i][j][k])
				if errCheck(err) {
					fmt.Println("Ошибка в вводе коэффициента ", k, " степени переменной x[", j, "]")
					os.Exit(1)
				}

			}
		}
	}
}

// init valuse input
func inputValues() {
	fmt.Print("### Ввод начальных значений ###\n")
	for j := 0; j < perN; j++ {
		fmt.Print("Начальное значение переменной x[", j, "]= ")
		_, err := fmt.Scan(&values[j])
		if errCheck(err) {
			fmt.Println("Ошибка в вводе начальных значений")
			os.Exit(1)
		}
	}
	// user bound input
	fmt.Print("### Ввод порога сходимости (epsilon) ###\n")
	_, err := fmt.Scan(&epsilon)
	if errCheck(err) {
		fmt.Println("Ошибка в вводе порога")
		os.Exit(1)
	}
}

// group input functions in big one
func input() {
	inputNumbers()
	inputFactors()
	inputValues()
}

// create derivative factors of every stroke
func derivative() {
	for i := 0; i < lineN; i++ {
		for j := 0; j < perN; j++ {
			for k := 0; k < maxP[i][j]; k++ {
				D[i][j][k] = F[i][j][k+1] * float64(k+1)
			}
		}
	}
}

// solving a polinom, derivative or not
func solve(res float64, x float64, mas [L][A][P]float64, i int, j int, power int) float64 {
	for k := power; k >= 0; k-- {
		res = res*x + mas[i][j][k]
	}
	return res
}

// calculating gradient for every assignment
func gradient(sum float64, x float64, j int) float64 {
	for i := 0; i < lineN; i++ {
		fmt.Println("................................х:", x)
		sum = sum + 2*solve(0, x, F, i, j, maxP[i][j])*solve(0, x, D, i, j, maxP[i][j]-1)
	}
	return sum
}

// final output Ф(x)
func builder(sum float64, bound float64) (result float64, resultMin float64) {
	for j := 0; j < perN; j++ {
		for i := 0; i < lineN; i++ {
			sum = sum + math.Pow(solve(0, values[j], F, i, j, maxP[i][j]), 2)
		}
	}
	fmt.Println("Ф(Х)= ", sum)
	if sum < bound {
		bound = sum
		for j := 0; j < perN; j++ {
			valuesMin[j] = values[j]
		}
	}
	return sum, bound
}

func main() {
	fmt.Println("\n\n\nРешение системы нелинейных уравнений методом градиентного спуска")
	fmt.Print("----------------------------------------------------------------\n\n")
	input()
	derivative()

	// step of descent
	step := 0.0001
	bound := 10000000000000.01

	var result, prevResult float64
	for {
		for j := 0; j < perN; j++ {
			grad[j] = gradient(0, values[j], j)
		}
		for j := 0; j < perN; j++ {
			values[j] = values[j] - step*grad[j]/float64(perN)
		}
		prevResult = result
		result, bound = builder(0, bound)
		// if step(n) - step(n-1) < epsilon: exit
		if math.Abs(result-prevResult) < epsilon {
			break
		}
		if result > bound {
			fmt.Println("Градиент не сходится")
			break
		}

	}
	// output result and x values
	fmt.Print("Ф(Х)= ", bound, "\n")
	for j := 0; j < perN; j++ {
		fmt.Println("x[", j, "]= ", valuesMin[j])
	}
}
