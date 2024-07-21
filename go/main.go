package main

import (
	"fmt"
	"sort"
)

type Humans struct {
	age  int
	name string
}
type HumanSlice []Humans

func (h HumanSlice) Len() int {
	return len(h)
}
func (h HumanSlice) Less(i, j int) bool {
	return h[i].age < h[j].age
}
func (h HumanSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[j]
}
func main() {
	data := HumanSlice{{age: 3, name: "tom"}, {age: 2, name: "jon"}, {age: 5, name: "pika"}, {age: 1, name: "pan"}}
	sort.Sort(data)
	fmt.Println(data)
}
