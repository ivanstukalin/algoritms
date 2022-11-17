package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	tests()
}

func tests() {
	// LuckyTicketsFor6DigitsNumber()
	// TestLuckyTickets("LuckyTickets")
	TestIterationExponentiation()
	TestFibonacci()
	TestPrimeNumbers()
}

func TestLuckyTickets(testName string) {
	luckyTickets := LuckyTickets{}
	for i := 0; i <= 9; i++ {
		N, _ := strconv.Atoi(readFromFile("test_data/" + testName + "/test." + strconv.Itoa(i) + ".in"))
		result := luckyTickets.Test(N)
		validValue, _ := strconv.Atoi(readFromFile("test_data/" + testName + "/test." + strconv.Itoa(i) + ".out"))

		fmt.Println(i, ": ", result == validValue)
	}
}

func TestIterationExponentiation() {
	e := Exponentiation{}
	fmt.Println("Возведение в степень: ", e.Test(2, 8) == 256)
	fmt.Println("Возведение в степень через степени двойки: ", e.TestExponentiationDegreeOfTwo(2, 9) == 512)
	fmt.Println("Возведение в степень log(N): ", e.TestExponentiationLogN(5, 20) == 95367431640625)
}

func TestFibonacci() {
	f := Fibonacci{}
	fmt.Println("Фибоначчи итеративный: ", f.Test(9) == 34)
	fmt.Println("Фибоначчи рекурсивный: ", f.TestFibonacciRecursive(8) == 21)
	fmt.Println("Фибоначчи золотое сечение: ", f.TestFibonacciGold(9) == 34)
	fmt.Println("Фибоначчи с помощью матриц: ", f.TestFibonacciMatrix(10) == 21+34)
}

func TestPrimeNumbers() {
	p := PrimeNumbers{}
	counts := []int{10, 100, 1000, 10000, 100000}

	fmt.Println("Простые числа O(n^2): ")
	for _, count := range counts {
		start := time.Now()
		fmt.Println(count, "->: ", p.TestPrimeNumbersIn(count))
		fmt.Println(time.Since(start))
	}
	counts = []int{10, 100, 1000, 10000, 100000, 1000000, 10000000}
	for _, count := range counts {
		start := time.Now()
		fmt.Println(count, "->: ", p.TestPrimeNumbersSqrt(count))
		fmt.Println(time.Since(start))
	}
	for _, count := range counts {
		start := time.Now()
		fmt.Println(count, "->: ", p.TestPrimeNumbersByPrimeNumbers(count))
		fmt.Println(time.Since(start))
	}
	for _, count := range counts {
		start := time.Now()
		fmt.Println(count, "->: ", p.TestPrimeNumbersEratosthenes(count))
		fmt.Println(time.Since(start))
	}
}

func readFromFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	return wr.String()
}
