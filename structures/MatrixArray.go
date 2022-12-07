package structures

type MatrixArrayInterface interface {
	PlusOne(value int)
	DeleteOne(index int) int
}

type MatrixArray struct {
	LastIndex    int
	Values       map[int]VectorArray
	ResizeCount  int
	CurrentArray VectorArray
}

func (ma *MatrixArray) PlusOne(value int) {
	if ma.LastIndex == ma.CurrentArray.Vector || len(ma.CurrentArray.Values) == 0 {
		ma.resize()
	}
	ma.CurrentArray.PlusOne(value)
	ma.LastIndex++
}
func (ma *MatrixArray) DeleteOne(index int) int {
	arrNumber := index / ma.CurrentArray.Vector
	valNumber := index % ma.CurrentArray.Vector

	ma.CurrentArray = ma.Values[arrNumber]
	var deletedValue int

	for i := arrNumber; i < len(ma.Values); i++ {
		for j := 0; j < ma.CurrentArray.Vector; j++ {
			if arrNumber == i && j <= valNumber {
				deletedValue = ma.Values[i].Values[j]
				continue
			}
			if j == 0 {
				ma.Values[i-1].Values[ma.CurrentArray.Vector-1] = ma.Values[i].Values[j]
				continue
			}
			ma.Values[i].Values[j-1] = ma.Values[i].Values[j]
		}
	}
	ma.Values[len(ma.Values)-1].Values[len(ma.CurrentArray.Values)-1] = 0
	return deletedValue
}

func (ma *MatrixArray) resize() {
	newArr := VectorArray{
		Vector: 5,
		Values: make([]int, 5),
	}
	next := len(ma.Values)
	if len(ma.Values) == 0 {
		ma.Values = map[int]VectorArray{}
	}
	ma.Values[next] = newArr
	ma.CurrentArray = newArr
	ma.LastIndex = 0
}

func (ma *MatrixArray) Drop() {
	newArr := map[int]VectorArray{}
	ma.Values = newArr
	ma.LastIndex = 0
}
