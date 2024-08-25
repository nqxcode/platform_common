package gob

import (
	"bytes"
	"encoding/gob"
)

// DeepClone deep clone
func DeepClone(src, dst any) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(src)
	if err != nil {
		panic(err)
	}

	err = dec.Decode(dst)
	if err != nil {
		panic(err)
	}
}
