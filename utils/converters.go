package utils

import (
	"bytes"
	"encoding/binary"
)

func Float64ToByte(value float64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func StringToByte(value string) []byte {
	return []byte(value)
}
