package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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
}

func TestFibonacci() {
	f := Fibonacci{}
	fmt.Println("Фибоначчи итеративный: ", f.Test(9) == 34)
	fmt.Println("Фибоначчи рекурсивный: ", f.TestFibonacciRecursive(8) == 21)
}

func TestPrimeNumbers() {
	p := PrimeNumbers{}
	fmt.Println("Простые числа O(n^2): ", p.TestPrimeNumbersIn(100) == 25)
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
