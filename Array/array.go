package array

import "errors";


type Array struct {
	values   []interface{}
	isStatic bool
}

func NewStaticArray(size int) *Array {
	return &Array{
		values:   make([]interface{}, size),
		isStatic: true,
	}
}
func NewDynamicArray() *Array {
	return &Array{
		isStatic: false,
	}
}

func (a *Array) Push(value interface{}) error {
	if a.isStatic  && len(a.values) >= cap(a.values) {
		return errors.New("Array capacity is full")
	}
	a.values = append(a.values, value)
	return nil
}
func (a *Array) Pop() interface{} {
	lastElement := a.values[len(a.values)-1]
	a.values = a.values[:len(a.values)-1]
	return lastElement
}
func (a *Array) ToArray() []interface{} {
	return a.values
}

func (a *Array) Len() int {
	return len(a.values)
}
