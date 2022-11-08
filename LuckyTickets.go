package main

import "fmt"

func main() {
	GetTicketsBy2NDigitsNumber(3)
}

func GetTicketsBy2NDigitsNumber(N int) {
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
	fmt.Println(result)
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

/// xxx0
/// 0 -> 0000
/// 1 -> 1000 0100 0010
/// 2 -> 1100 0110 0020 2000 0200 1010
///
