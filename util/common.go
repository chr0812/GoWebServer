package util

import "encoding/json"

func Marshal(v interface{}) []byte {
	rt, _ := json.Marshal(v)
	return rt
}
