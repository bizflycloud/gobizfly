package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MergeQueries(req *http.Request, paramStr interface{}) error {
	var params map[string]interface{}
	if req == nil || paramStr == nil {
		return nil
	}
	paramBytes, err := json.Marshal(paramStr)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(paramBytes, &params); err != nil {
		return err
	}

	q := req.URL.Query()
	for key, val := range params {
		valStr := fmt.Sprintf("%v", val)
		q.Add(key, valStr)
	}
	req.URL.RawQuery = q.Encode()
	return nil
}
