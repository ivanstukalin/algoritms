package main

import (
	"math"
)

type FibonacciInterface interface {
	Test(num int) int
	TestFibonacciRecursive(num int) int
	TestFibonacciGold(num int) int
	TestFibonacciMatrix(num int) int
}

type Fibonacci struct {
}

var Fi float64 = (1 + math.Sqrt(5)) / 2

func fibonacci(num int) int {
	fibonacciNumbers := []int{0, 1}
	if num < 3 {
		return fibonacciNumbers[num-1]
	}
	for i := 2; len(fibonacciNumbers) <= num; i++ {
		fibonacciNumbers = append(fibonacciNumbers, fibonacciNumbers[len(fibonacciNumbers)-1]+fibonacciNumbers[len(fibonacciNumbers)-2])
	}
	return fibonacciNumbers[len(fibonacciNumbers)-1]
}

func fibonacciRecursive(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fibonacciRecursive(num-1) + fibonacciRecursive(num-2)
}

func fibonacciGold(num int) int {
	result := math.Pow(Fi, float64(num))/math.Sqrt(5) + 1/2
	return int(Round(result))
}

func fibonacciMatrix(num int) int {
	baseMarix := [][]int{
		{1, 1},
		{1, 0},
	}
	e := Exponentiation{}
	finalMatrix := e.MatrixExponentiationLogN(baseMarix, num-1)
	return finalMatrix[0][0]
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func (f *Fibonacci) Test(num int) int {
	return fibonacci(num)
}

func (f *Fibonacci) TestFibonacciRecursive(num int) int {
	return fibonacciRecursive(num)
}

func (f *Fibonacci) TestFibonacciGold(num int) int {
	return fibonacciGold(num)
}

func (f *Fibonacci) TestFibonacciMatrix(num int) int {
	return fibonacciMatrix(num)
}
