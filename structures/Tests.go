package structures

import (
	"fmt"
	"time"

	"github.com/emirpasic/gods/lists/arraylist"
)

func Tests() {
	testSingleArray()
	testVectorArray()
	testFactorArray()
	testMatrixArray()
	testArrayList()
}

func testSingleArray() {
	fmt.Println("Single:")
	singleArray := SingleArray{}
	testPlusOne(&singleArray)
	testDeleteOne(&singleArray, 53)
}

func testVectorArray() {
	fmt.Println("Vector:")
	vectorArray := VectorArray{
		Vector: 2,
	}
	testPlusOne(&vectorArray)
	testDeleteOne(&vectorArray, 53)
}

func testFactorArray() {
	fmt.Println("Factor:")
	factorArray := FactorArray{
		Factor: 10,
	}
	testPlusOne(&factorArray)
	testDeleteOne(&factorArray, 53)
}

func testMatrixArray() {
	fmt.Println("Matrix:")
	matrixArray := MatrixArray{}
	testPlusOne(&matrixArray)
	testDeleteOne(&matrixArray, 53)
}

func testArrayList() {
	list := arraylist.New([]int{})
	tests := []int{10, 100, 1000, 10000, 100000}
	for _, test := range tests {
		start := time.Now()
		for i := 0; i < test; i++ {
			list.Add(i)
		}
		fmt.Println("Добавление ", test, " элементов -", time.Since(start))
		list.Clear()
	}

	for i := 0; i < 100; i++ {
		list.Add(i)
	}
	start := time.Now()
	list.Remove(53)
	fmt.Println("Удаленное значение: ", 53, ", время выполнения: ", time.Since(start))
}

func testPlusOne(array ArrayInterface) {
	tests := []int{10, 100, 1000, 10000, 100000}
	for _, test := range tests {
		start := time.Now()
		for i := 0; i < test; i++ {
			array.PlusOne(i)
		}
		fmt.Println("Добавление ", test, " элементов -", time.Since(start))
		array.Drop()
	}
}

func testDeleteOne(array ArrayInterface, index int) {
	for i := 0; i < 100; i++ {
		array.PlusOne(i)
	}
	start := time.Now()
	result := array.DeleteOne(index)
	fmt.Println("Удаленное значение: ", result, ", время выполнения: ", time.Since(start))
}
