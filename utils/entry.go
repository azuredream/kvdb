package utils

import (
	"encoding/binary",
	"time",
)

type ValueStruct struct {
	Meta byte
	Value []byte
	TTL uint64
	Version uint64
}

func (vs *ValueStruct) EncodeSize() uint32{
	size := len(vs.Value) + 1
	enc :=sizeVarint(vs.ExpiresAt)
	return uint32(sz+enc)
}
// Meta | ExpiresAt | value
func (vs *ValueStruct) EncodeValue(b []byte) uint32{
	b[0] = vs.Meta
	sz := binary.PutUvarint(b[1:], vs.ExpiresAt)
	n := copy(b[1+sz:], vs.Value)
}

//Byte length of TTL
func sizeVarint(x uint64) (n int) {
	for {
		n++
		x>>=7
		if x == 0 {
			break
		}
	}
	return n
}