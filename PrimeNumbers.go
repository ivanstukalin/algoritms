package main

type PrimeNumbersInterface interface {
	TestPrimeNumbersIn(num int) int
}

type PrimeNumbers struct {
}

func primeNumbersIn(num int) int {
	result := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			result++
		}
	}
	return result
}

func isPrime(num int) bool {
	for i := 2; i <= num-1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (p *PrimeNumbers) TestPrimeNumbersIn(num int) int {
	return primeNumbersIn(num)
}
