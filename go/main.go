package main

import (
	"fmt"
	"sort"
)

type Human struct {
	name string
	age  int
}
type Interface interface {
	// Find number of elements in collection
	Len() int

	// Less method is used for identifying which elements among index i and j are lesser and is used for sorting
	Less(i, j int) bool

	// Swap method is used for swapping elements with indexes i and j
	Swap(i, j int)
}
type AgeFactor []Human

func (a AgeFactor) Len() int           { return len(a) }
func (a AgeFactor) Less(i, j int) bool { return a[i].age < a[j].age }
func (a AgeFactor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}
	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)
}
