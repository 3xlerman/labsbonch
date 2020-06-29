package main

import (
	"fmt"
	"math"
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
