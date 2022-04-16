package util

import "encoding/json"

func UnpackRequest(request []byte, o interface{}) error {
	if err := json.Unmarshal(request, &o); nil != err {
		return err
	}
	return nil
}
