package main_test

import (
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
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
	initValues [A]float64

	// array of gradient values
	grad [A]float64

	// epsilon
	epsilon float64

	// G gradient of all assignments value
	G float64

	bound float64 = 10000000000000.01

	isQ int

	tMin float64 = 1.128

	out outputData
)

type outputData struct {
	ResultValues []float64
	FuncVal      float64
	Iterations   int
	Time         time.Duration
}

func BenchmarkSample1(b *testing.B) {

	lineN = 1
	perN = 1

	maxP[0][0] = 2

	F[0][0][2] = 1
	F[0][0][1] = 2
	F[0][0][0] = 10

	values[0] = 0

	epsilon = 0.1

	outputV := make([]float64, perN)

	for p := 0; p < b.N; p++ {
		derivative()
		// step of descent
		step := 0.0001

		iter := 1
		for {
			for j := 0; j < perN; j++ {
				grad[j] = gradient(0, values[j], j)
			}
			G = calculateG(0, grad)
			if G < epsilon {
				break
			}
			for j := 0; j < perN; j++ {
				values[j] = values[j] - step*grad[j]/float64(perN)
			}
			result, bound := builder(0)
			if result > bound {
				break
			}

			iter++
		}


		for j := 0; j < perN; j++ {
			outputV[j] = values[j]
		}
	}
}

func BenchmarkSample2(b *testing.B) {
	lineN = 1
	perN = 1

	maxP[0][0] = 2

	F[0][0][2] = 1
	F[0][0][1] = 2
	F[0][0][0] = 10

	values[0] = 0

	epsilon = 0.1

	outputV := make([]float64, perN)

	for p := 0; p < b.N; p++ {

		isQ = 0
		iter := 0
		for j := 0; j < perN; j++ {
			values[j] = initValues[j]
		}

		for {
			if iter == 0 {
				for j := 0; j < perN; j++ {
					grad[j] = gradient(0, values[j], j)
				}

				G = calculateG(0, grad)

				if G < epsilon {
					break
				}

				t := findStep(values, grad)

				for j := 0; j < perN; j++ {
					values[j] = values[j] - t*G
				}

				iter++

			} else {
				gradPrev := grad
				for j := 0; j < perN; j++ {
					grad[j] = gradient(0, values[j], j)
				}
				GPrev := G
				G = calculateG(0, grad)

				if G < epsilon {
					break
				}
				var b float64
				if isQ == 1 {
					b = G*G / GPrev*GPrev
				} else {
					z := 0.0
					c := 0.0
					for j := 0; j < perN; j++ {
						z = z + grad[j]*(grad[j]-gradPrev[j])
						c = c + grad[j]*grad[j]
					}
					b = z / c
				}

				d := G + b*GPrev

				t := findStep(values, grad)

				for j := 0; j < perN; j++ {
					values[j] = values[j] - t*d
				}

				iter++
			}

		}
		_, bound = builder(0)


		for j := 0; j < perN; j++ {
			outputV[j] = values[j]
		}
	}
}

func gradientDescent(c chan outputData, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Градиентный спуск")
	derivative()
	// step of descent
	step := 0.0001

	iter := 1
	t0 := time.Now()
	for {
		for j := 0; j < perN; j++ {
			grad[j] = gradient(0, values[j], j)
		}
		G = calculateG(0, grad)
		if G < epsilon {
			break
		}
		for j := 0; j < perN; j++ {
			values[j] = values[j] - step*grad[j]/float64(perN)
		}
		result, bound := builder(0)
		if result > bound {
			fmt.Println("Градиент не сходится")
			break
		}

		iter++
	}

	outputV := make([]float64, perN)
	for j := 0; j < perN; j++ {
		outputV[j] = values[j]
	}

	// output result and x values
	out.FuncVal = bound
	out.ResultValues = outputV
	out.Iterations = iter
	out.Time = time.Since(t0)
	deleteValues()
	c <- out

}

func conjugateGradients(c chan outputData, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Сопряжённые градиенты")
	isQ = 0
	iter := 0
	for j := 0; j < perN; j++ {
		values[j] = initValues[j]
	}
	t0 := time.Now()
	for {
		if iter == 0 {
			for j := 0; j < perN; j++ {
				grad[j] = gradient(0, values[j], j)
			}

			G = calculateG(0, grad)

			if G < epsilon {
				break
			}

			t := findStep(values, grad)

			for j := 0; j < perN; j++ {
				values[j] = values[j] - t*G
			}

			iter++

		} else {
			gradPrev := grad
			for j := 0; j < perN; j++ {
				grad[j] = gradient(0, values[j], j)
			}
			GPrev := G
			G = calculateG(0, grad)

			if G < epsilon {
				break
			}
			var b float64
			if isQ == 1 {
				b = math.Pow(G, 2) / math.Pow(GPrev, 2)
			} else {
				z := 0.0
				c := 0.0
				for j := 0; j < perN; j++ {
					z = z + grad[j]*(grad[j]-gradPrev[j])
					c = c + grad[j]*grad[j]
				}
				b = z / c
			}

			d := G + b*GPrev

			t := findStep(values, grad)

			for j := 0; j < perN; j++ {
				values[j] = values[j] - t*d
			}

			iter++
		}

	}
	_, bound = builder(0)

	outputV := make([]float64, perN)
	for j := 0; j < perN; j++ {
		outputV[j] = values[j]
	}

	out.FuncVal = bound
	out.ResultValues = outputV
	out.Iterations = iter
	out.Time = time.Since(t0)
	c <- out

}

func findStep(values [A]float64, grad [A]float64) float64 {
	min := 1e+308
	for j := 0; j < perN; j++ {
		for i := 0; i < lineN; i++ {
			g := 0.0
			for t := 0.001; t < 1000; t *= 1.1 {

				g = math.Abs(2 * solve(0, values[j]-t*grad[j], F, i, j, maxP[i][j]) * derivativeValue(0, values[j]-t*grad[j], F, maxP[i][j], i, j))
				if g < min {
					tMin = t
					break
				}
			}
		}
	}
	return tMin
}

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
		initValues[j] = values[j]
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

// create derivative factors of every stroke
func derivativeValue(sum float64, x float64, mas [L][A][P]float64, power int, i int, j int) float64 {
	for i := 0; i < lineN; i++ {
		for j := 0; j < perN; j++ {
			for k := power; k > 0; k-- {
				sum = sum + float64(k)*mas[i][j][k]*math.Pow(x, float64(power-1))
			}
		}
	}
	return sum
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
		sum = sum + 2*solve(0, x, F, i, j, maxP[i][j])*solve(0, x, D, i, j, maxP[i][j]-1)
	}
	return sum
}

// final output Ф(x)
func builder(sum float64) (result float64, resultMin float64) {
	for j := 0; j < perN; j++ {
		for i := 0; i < lineN; i++ {
			sum = sum + math.Pow(solve(0, values[j], F, i, j, maxP[i][j]), 2)
		}
	}
	if sum < bound {
		bound = sum
	}
	return sum, bound
}

func calculateG(sum float64, grad [A]float64) float64 {
	for j := 0; j < perN; j++ {
		sum = sum + grad[j]*grad[j]
	}
	return math.Sqrt(sum)
}

func deleteValues() {

	for j := 0; j < perN; j++ {
		values[j] = 0
		grad[j] = 0
	}

	bound = 10000000000000.01
}
