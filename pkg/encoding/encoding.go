package encoding

import (
	"bytes"
	"encoding/json"
)

func UnPackJSON(in interface{}, target interface{}) error {
	var e1 error
	var b []byte
	switch in := in.(type) {
	case []byte:
		b = in

	default:
		b, e1 = json.Marshal(in)
		if e1 != nil {
			return e1
		}
	}

	buf := bytes.NewBuffer(b)
	enc := json.NewDecoder(buf)
	enc.UseNumber()
	if err := enc.Decode(&target); err != nil {
		return err
	}
	return nil
}
