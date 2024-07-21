package exercises

import (
	"fmt"
	"strings"
)

func Exercise4() {
	res := summer("I", "am", "rahat")
	fmt.Println("Exercise4: ", res)
}
func summer(a ...string) (sum string) {
	sum = strings.Join(a, " ")
	return
}
