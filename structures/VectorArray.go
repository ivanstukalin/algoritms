package structures

type VectorArrayInterface interface {
	PlusOne(value int)
	DeleteOne(index int) int
	Drop()
}

type VectorArray struct {
	LastIndex   int
	Values      []int
	Vector      int
	ResizeCount int
}

func (va *VectorArray) PlusOne(value int) {
	if len(va.Values) == va.LastIndex {
		va.resize()
		va.ResizeCount++
	}
	va.Values[va.LastIndex] = value
	va.LastIndex++
}

func (va *VectorArray) DeleteOne(index int) int {
	newSize := len(va.Values) - 1
	var newArr = make([]int, newSize)
	var deletedValue int
	var counter int

	for i := 0; i < newSize; i++ {
		if i == index {
			deletedValue = va.Values[i]
			continue
		}
		newArr[counter] = va.Values[i]
		counter++
	}

	va.Values = newArr
	return deletedValue
}

func (va *VectorArray) resize() {
	var newArr = make([]int, len(va.Values)+va.Vector)

	for key, value := range va.Values {
		newArr[key] = value
	}

	va.Values = newArr
}

func (va *VectorArray) Drop() {
	newArr := []int{}
	va.Values = newArr
	va.LastIndex = 0
}
