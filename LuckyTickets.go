package main

type LuckyTicketsInterface interface {
	Test(N int) int
}

type LuckyTickets struct {
}

func (l *LuckyTickets) GetTicketsBy2NDigitsNumber(N int) int {
	baseArr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var arr []int
	arr = baseArr
	for i := 0; i < N-1; i++ {
		arr = nextArr(arr)
	}
	var result int
	for _, possibleSum := range arr {
		result += possibleSum * possibleSum
	}
	return result
}

func nextArr(prevArr []int) []int {
	nextArrLen := len(prevArr) + 9
	var nextArr []int
	for i := 0; i < nextArrLen; i++ {
		currentValue := 0
		for j := 0; j < 10; j++ {
			prevValuePos := i - j
			if prevValuePos >= 0 && prevValuePos <= len(prevArr)-1 {
				currentValue += prevArr[i-j]
			}
		}
		nextArr = append(nextArr, currentValue)
	}
	return nextArr
}

func (l *LuckyTickets) Test(N int) int {
	return l.GetTicketsBy2NDigitsNumber(N)
}
