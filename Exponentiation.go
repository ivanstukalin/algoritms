package main

type ExponentiationInterface interface {
	Test(num int, degree int) int
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

func (e *Exponentiation) Test(num int, degree int) int {
	return exponentiation(num, degree)
}
