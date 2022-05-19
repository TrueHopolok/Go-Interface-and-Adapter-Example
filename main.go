package main

import (
	"fmt"
	"go_interfaces/sortable"
)

type SortableInt int

func (f SortableInt) Before(s sortable.Sortable) bool {
	return f < SortableInt(s)
}

func main() {
	arr := []SortableInt{2, 5, 2, 9, 1, 8, 5, 3}
	fmt.Println(arr)
	sortable.MergeSort(arr)
	fmt.Printf("%t | %v", arr, arr)
}
