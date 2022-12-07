package structures

type FactorArrayInterface interface {
	PlusOne(value int)
	DeleteOne(index int) int
}

type FactorArray struct {
	LastIndex   int
	Values      []int
	Factor      int
	ResizeCount int
}

func (fa *FactorArray) PlusOne(value int) {
	if len(fa.Values) == fa.LastIndex {
		fa.resize()
		fa.ResizeCount++
	}
	fa.Values[fa.LastIndex] = value
	fa.LastIndex++
}
func (fa *FactorArray) DeleteOne(index int) int {
	newSize := len(fa.Values) - 1
	var newArr = make([]int, newSize)
	var deletedValue int
	var counter int

	for i := 0; i < newSize; i++ {
		if i == index {
			deletedValue = fa.Values[i]
			continue
		}
		newArr[counter] = fa.Values[i]
		counter++
	}

	fa.Values = newArr
	return deletedValue
}

func (fa *FactorArray) resize() {
	increament := len(fa.Values) * fa.Factor / 100
	if increament == 0 {
		increament = 1
	}

	var newArr = make([]int, len(fa.Values)+increament)

	for key, value := range fa.Values {
		newArr[key] = value
	}

	fa.Values = newArr
}

func (fa *FactorArray) Drop() {
	newArr := []int{}
	fa.Values = newArr
	fa.LastIndex = 0
}
