package utils

import (
	"bytes"
	"encoding/gob"
	"pepper/runtime"
)

func EncodeBytecode(instrs []runtime.VMInstr) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(instrs)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBytecode(data []byte) ([]runtime.VMInstr, error) {
	var instrs []runtime.VMInstr
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&instrs)
	if err != nil {
		return nil, err
	}
	return instrs, nil
}
