/*This skiplist is a mutable kv table in memory. 
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
	_ "unsafe"
)

const (
	MAXHEIGHT      = 20
	heightIncrease = math.MaxUint32 / 3
)

type node struct {
	v uint64 //
	koffset uint32,
	ksize uint32,
	height uint32,
	tower [MAXHEIGHT]uint32,


}

type skiplist struct {
	arena *Arena,
	head uint32,
	h uint32,
	ref uint32, //?
}

