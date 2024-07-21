package exercises

import (
	"fmt"
)

func Exercise3() {
	map_obj := make(map[string]struct{})
	map_obj["apple"] = struct{}{}
	res, ok := map_obj["cat"]
	fmt.Println("Exercise3:  ", map_obj, res, ok)
	//copyMap := make(map[string]struct{})
	//copy(map_obj,copyMap)

}
