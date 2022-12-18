/*
This skiplist is a mutable kv table in memory.
This table will be transformed to immutable table and then SStable.

Use cases(APIs)
node

	init(key,value)
	setvalue(value)
	getvalue()
	getkey()
	getnext(level) //atomic load
	setnext(oldoffset, newoffset, level)//atomic cas

encodevalue(voffset,vsize)
decodevalue(value)
skiplist

	init

	getvalue(key)
	putkv(key,value)
	setvalue(key,value)

Define Class(data + action) and relationship

	const(maxheight = 30)
	class node{
		koffset,
		ksize
		voffset,
		vsize,
		height,
		tower []uint32,
	}
	class skiplist{
		arena Arena,
		head,
		height,
	}
*/
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

type skiplist struct {
	arena *Arena,
	headoffset uint32,
	h uint32,
	ref uint32, //?
}


// type skiplist struct {
// 	arena *Arena,
// 	headoffset uint32,
// 	h uint32,
// 	ref uint32, //?
// }
func newskiplist (arenasize uint64) *skiplist {
	arena := newArena(arenasize)
	//get head
	nodesize := MAXNODE
	head  := newnode(arena, nil, nil, MAXHEIGHT)
	headoffset := uint32(uintprt(unsafe.Pointer(head)) - uintprt(unsafe.Pointer(arena.buf[0])))
	return &skiplist{
		arena: arena,
		headoffset: headoffset,
		h: 1,
		ref: 0,
	}

}

// type node struct {
// 	v uint64,
// 	koffset uint32,
// 	ksize uint32,
// 	height uint32,
// 	tower [MAXHEIGHT]uint32,
// }

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

func (n* node) encodev(vsize uint32, voffset uint32) uint64 {
	return uint64(vsize<<32) | uint64(voffset)
}

func (n* node) decodev(codedv uint64) vsize uint32, voffset uint32 {
	vsize = uint32(codedv>>32)
	voffset = uint32(codedv)
	return
}