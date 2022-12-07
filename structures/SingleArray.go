package structures

type SingleArrayInterface interface {
	PlusOne(value int)
	DeleteOne(index int) int
}

type SingleArray struct {
	Size   int
	Values []int
}

func (sa *SingleArray) PlusOne(value int) {
	sa.Size = len(sa.Values)
	sa.resize()
	sa.Values[sa.Size] = value
}

func (sa *SingleArray) DeleteOne(index int) int {
	newSize := sa.Size - 1
	var newArr = make([]int, newSize)
	var deletedValue int
	var counter int

	for i := 0; i < newSize; i++ {
		if i == index {
			deletedValue = sa.Values[i]
			continue
		}
		newArr[counter] = sa.Values[i]
		counter++
	}

	sa.Values = newArr
	return deletedValue
}

func (sa *SingleArray) resize() {
	newSize := sa.Size + 1
	var newArr = make([]int, newSize)

	for key, value := range sa.Values {
		newArr[key] = value
	}

	sa.Values = newArr
}

func (sa *SingleArray) Drop() {
	newArr := []int{}
	sa.Values = newArr
}