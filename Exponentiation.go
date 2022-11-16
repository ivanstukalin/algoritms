package main

import (
	"math"
)

type ExponentiationInterface interface {
	Test(num int, degree int) int
	TestExponentiationDegreeOfTwo(num int, degree int) int
	TestExponentiationLogN(num int, degree int) int
}

type Exponentiation struct {
}

func exponentiation(num int, degree int) int {
	result := num
	for i := 2; i <= degree; i++ {
		result = result * num
	}
	return result
}

func exponentiationDegreeOfTwo(num int, degree int) int {
	result := num
	degreeOfTwo := int(math.Log2(float64(degree)))
	for i := 0; i < degreeOfTwo; i++ {
		result = result * result
	}
	last := degree - int(math.Pow(2, float64(degreeOfTwo)))
	if last > 0 {
		result *= exponentiation(num, last)
	}
	return result
}

func exponentiationLogN(num int, degree int) int {
	result := 1
	temp := num
	for degree > 1 {
		degree = int(degree / 2)
		temp = temp * temp
		if degree%2 != 0 {
			result *= temp
		}
	}
	return result
}

func (e *Exponentiation) Test(num int, degree int) int {
	return exponentiation(num, degree)
}

func (e *Exponentiation) TestExponentiationDegreeOfTwo(num int, degree int) int {
	return exponentiationDegreeOfTwo(num, degree)
}

func (e *Exponentiation) TestExponentiationLogN(num int, degree int) int {
	return exponentiationLogN(num, degree)
}
