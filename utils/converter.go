package utils

import "encoding/json"

func ConvDataWithJson[T any](data interface{}) (T, error) {
	var result T
	byteData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(byteData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
