package util

import (
	"math"
	"unsafe"
	_ "unsafe"
)

const (
	MAXHEIGHT      = 20
	heightIncrease = math.MaxUint32 / 3
	MAXNODE = unsafe.Sizeof(node{})
)

type node struct {
	v uint64,
	koffset uint32,
	ksize uint32,
	height uint32,
	tower [MAXHEIGHT]uint32,
}


func newnode (a arena, k []byte, v []byte, height uint32) *node {
	//set up node meta data
	nodeoffset = a.allocate(MAXNODE);
	keysize := uint32(len(k))
	koffset := a.allocate(keysize)
	vsize := uint32(len(v))
	voffset := a.allocate(vsize)
	node := (*node)unsafe.Pointer(&s.buf[koffset])
	node.keysize = keysize
	node.koffset = koffset
	node.v = encodev(vsize, voffset)
	node.height = height
	//setupk
	AssertTrue(len(k) == copy(a.buf[koffset:koffset+ksize], k))
	//setup v
	AssertTrue(len(v) == copy(a.buf[voffset:voffset+vsize], v))
	//
	return node

}





