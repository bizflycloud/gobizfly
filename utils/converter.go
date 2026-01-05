package utils

import "encoding/json"

func ConvDataWithJson(data interface{}, result interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteData, &result)
}
