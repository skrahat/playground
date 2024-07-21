package exercises

import (
	"fmt"
)

func Exercise4() {
	res := summer(2, 3, 4, 5)
	fmt.Println("Exercise4:  %v", res)
}
func summer(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}
