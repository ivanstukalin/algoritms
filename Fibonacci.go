package main

type FibonacciInterface interface {
	Test(num int) int
	TestFibonacciRecursive(num int) int
}

type Fibonacci struct {
}

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

func (f *Fibonacci) Test(num int) int {
	return fibonacci(num)
}

func (f *Fibonacci) TestFibonacciRecursive(num int) int {
	return fibonacciRecursive(num)
}
