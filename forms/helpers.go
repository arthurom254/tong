package forms

import (
	"encoding/json"
	"fmt"
)

func StructToMap(in interface{}) {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(in)
	json.Unmarshal(inrec, &inInterface)

	// iterate through inrecs
	for field, val := range inInterface {
		fmt.Println("KV Pair: ", field, val)
	}
}
