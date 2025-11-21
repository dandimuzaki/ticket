package utils

import (
	"encoding/json"
	"fmt"
)

func StringToMap(jsonString string) (map[string]float64, error) {
	var data map[string]float64

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}
	return data, nil
}