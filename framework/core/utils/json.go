package utils

import "encoding/json"

func ToJsonStr(data interface{}) string {
	b, e := json.Marshal(data)
	if e != nil {
		return ""
	}
	return string(b)
}
