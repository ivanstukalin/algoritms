package main

import (
	"math"
)

type PrimeNumbersInterface interface {
	TestPrimeNumbersIn(num int) int
	TestPrimeNumbersByPrimeNumbers(num int) int
	TestPrimeNumbersEratosthenes(num int) int
}

type PrimeNumbers struct {
	PrimeNumbers []int
}

func primeNumbersIn(num int, algortim int) int {
	result := 0
	switch algortim {
	case 1:
		for i := 2; i <= num; i++ {
			if isPrime(i) {
				result++
			}
		}
	case 2:
		for i := 2; i <= num; i++ {
			if isPrimeSqrt(i) {
				result++
			}
		}
	}
	return result
}

func (p *PrimeNumbers) primeNumbersInWithCheckingOnlyPrimeNumbers(num int) int {
	p.PrimeNumbers = []int{}
	for i := 2; i <= num; i++ {
		if p.isPrimeByPrimeNumbers(i) {
			p.PrimeNumbers = append(p.PrimeNumbers, i)
		}
	}
	return len(p.PrimeNumbers)
}

func isPrime(num int) bool {
	for i := 2; i <= num-1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeSqrt(num int) bool {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (p *PrimeNumbers) isPrimeByPrimeNumbers(num int) bool {
	for _, primeNumber := range p.PrimeNumbers {
		if primeNumber > int(math.Sqrt(float64(num))) {
			return true
		}
		if num%primeNumber == 0 {
			return false
		}
	}
	return true
}

func primeNumbersEratosthenes(num int) int {
	numbers := make([]int, num)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if numbers[i] == 0 {
			for j := i * i; j < num; j += i {
				numbers[j] = 1
			}
		}
	}

	count := 0
	for i := 2; i < num; i++ {
		if numbers[i] == 0 {
			count++
		}
	}

	return count
}

func (p *PrimeNumbers) TestPrimeNumbersIn(num int) int {
	return primeNumbersIn(num, 1)
}

func (p *PrimeNumbers) TestPrimeNumbersSqrt(num int) int {
	return primeNumbersIn(num, 2)
}

func (p *PrimeNumbers) TestPrimeNumbersByPrimeNumbers(num int) int {
	return p.primeNumbersInWithCheckingOnlyPrimeNumbers(num)
}

func (p *PrimeNumbers) TestPrimeNumbersEratosthenes(num int) int {
	return primeNumbersEratosthenes(num)
}
