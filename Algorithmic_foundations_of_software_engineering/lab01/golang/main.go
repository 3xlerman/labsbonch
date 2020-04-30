package main

import (
	"fmt"
	"math"
)

const P = 100

type line struct { // объявляем структуру с данными о каждой строке в системе
	Power [P]int // максимальная степень каждой переменной x(0)-x(n)
	X     [P][100]float64 // массив коэффициентов для каждой переменной вплоть до 100-й степени
	D     [P][100]float64 // массив коэфициентов производной для каждой переменной
}

var (
	xvalue      [P]float64 // входные значения Х
	stroke      [P]line // поддержка до P строк в системе
	n, iter     int // кол-во строк и кол-во итераций
	grad        [P]float64 // массив для хранения градиентов для каждой переменной
	step, bound float64 // шаг градиентного спуска и поиск наиболее точного значения
	xmin        [P]float64 // хранение аргументов, при которых достигнуто наиболее точное значение функции
)

func input() { // функция ввода коэффициентов системы
	fmt.Print("Количество строк в системе: ")
	fmt.Scan(&n)
	for str := 0; str < n; str++ {
		fmt.Print("Строка ", str+1, "\n")
		for per := 0; per < n; per++ {
			fmt.Print("Переменная x[", per, "]\n")
			fmt.Print("Высшая степень переменной x[", per, "]: ")
			fmt.Scan(&stroke[str].Power[per])
			for k := stroke[str].Power[per]; k >= 0; k-- {
				fmt.Print("Введите коэффициент степени ", k, " = ")
				fmt.Scan(&stroke[str].X[per][k])
			}
		}
	}
}
func derivative() { // функция, вычисляющая коэффициенты в функции производной
	for str := 0; str < n; str++ {
		for per := 0; per < n; per++ {
			for k := 0; k < stroke[str].Power[per]; k++ {
				stroke[str].D[per][k] = stroke[str].X[per][k+1] * float64(k+1) // сохраняем коэффициенты в отдельный массив
			}
		}
	}
}
func solvePolinom(x float64, s [P][100]float64, per int, hpower int) float64 { // функция, находящая значение полинома при входном значении Х
	res := 0.0
	for k := hpower; k >= 0; k-- {
		res = res*x + s[per][k]
	}
	return res
}
func output(xvalue [P]float64, n int) { // функция вывода значения Ф(Х)
	sum := 0.0
	for per := 0; per < n; per++ {
		for str := 0; str < n; str++ {
			sum += math.Pow(solvePolinom(xvalue[per], stroke[str].X, per, stroke[str].Power[per]), 2)
		}
	}
	fmt.Print("Ф(Х)= ")
	fmt.Printf("%.5f\n", sum)
	if sum < bound { // если значение ниже минимального ранее
		for per := 0; per < n; per++ {
			xmin[per] = xvalue[per] // также запоминаем аргументы, при которых находим такое значение
		}
		bound = sum // данное значение становится минимальным
	}

}
func main() {
	bound = 100 // задаём минимум достаточно большим
	input() // ввод данных
	derivative() // нахождение коэффициентов производных для каждой переменной в каждой строке
	for per := 0; per < n; per++ {
		fmt.Print("Начальное значение переменной x[", per, "]: ")
		fmt.Scan(&xvalue[per])
	}

	fmt.Print("Введите шаг градиентного спуска: ")
	fmt.Scan(&step)
	fmt.Print("Введите количество итераций: ")
	fmt.Scan(&iter)
	for i := 0; i < iter; i++ {
		for per := 0; per < n; per++ {
			for str := 0; str < n; str++ { // нахоидим градиент по каждой переменной в Ф(Х)
				grad[per] = grad[per] + 2*solvePolinom(xvalue[per], stroke[str].X, per, stroke[str].Power[per])*solvePolinom(xvalue[per], stroke[str].D, per, stroke[str].Power[per]-1)
			}
			xvalue[per] = xvalue[per] - step*grad[per] // спускаемся
		}
		output(xvalue, n) // выводим значение Ф(Х)
	}
	fmt.Print("Наиболее точное решение Ф(Х)= ")
	fmt.Printf("%.5f%s\n", bound, " при значениях Х: ")
	for per := 0; per < n; per++ { // выводим все аргументы в системе
		fmt.Println("x[", per, "]= ", xmin[per])
	}
}
