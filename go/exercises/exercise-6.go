package exercises

import (
	"encoding/json"
	"log"
)

func Exercise6() {
	jsonStr := `{
		"Reading": ["John"],
		"Gaming": ["John", "Charlie"],
		"Cooking": ["Alice"],
		"Hiking": ["Alice", "Charlie"]
	}`

	var dataUnMarshal interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dataUnMarshal); err != nil {
		log.Fatalln("dataUnMarshal: ", err)
	}

	dataMap, ok := dataUnMarshal.(map[string]interface{})
	if !ok {
		log.Fatalln("Error: data is not a map")
	}
	result := make(map[string][]string)
	var count int
	// Iterate over the map
	for key, value := range dataMap {
		// Assert the type of the value to a slice of interfaces
		valueSlice, ok := value.([]interface{})
		if !ok {
			log.Println("Error: value is not a slice for key:", key)
			continue
		}

		// Print the key and its values
		//log.Printf("%s: ", key)
		for _, item := range valueSlice {
			// Assert the type of each item to a string
			str, ok := item.(string)
			if !ok {
				log.Println("Error: item is not a string")
				continue
			}
			//log.Printf("%s ", str)
			result[key] = append(result[key], str)
		}
		//log.Println()
		count++
	}
	log.Println("Exercise6: ", result)
}
