package main

type MatrixInterface interface {
	MatrixMultiplication2x2(matrix1 [][]int, matrix2 [][]int) [][]int
}

type Matrix struct {
}

func (m *Matrix) MatrixMultiplication2x2(matrix1 [][]int, matrix2 [][]int) [][]int {

	firstRow := []int{
		matrix1[0][0]*matrix2[0][0] + matrix1[0][1]*matrix2[1][0],
		matrix1[0][0]*matrix2[0][1] + matrix1[0][1]*matrix2[1][1],
	}
	secondRow := []int{
		matrix1[1][0]*matrix2[0][0] + matrix1[1][1]*matrix2[1][0],
		matrix1[1][0]*matrix2[0][1] + matrix1[1][1]*matrix2[1][1],
	}
	return [][]int{firstRow, secondRow}
}
