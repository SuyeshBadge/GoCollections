package main

import (
	"fmt"
	array "goCollections/Array"
)

func main() {
	myArray := array.NewStaticArray(5)
	fmt.Println(myArray)

	myArray.Push(10)
	myArray.Push(1)
	myArray.Push(0)
	myArray.Push(310)
	myArray.Push(310)
	myArray.Push(310)
	fmt.Print(myArray.ToArray())


}