/*
Arena is a memory pool
use cases:
priority:

	init,
	allocate(size)(offset uint32),
	size() uint32,
	capacity() uint32,

Define Class(data+operations) and Relationship:
//lockfree design

	class arena{
		capacity:
		cur:
		buf:
		<APIs>
	}

optimization: memory alignment
*/
package utils

import (
	"errors"
	"log"
	"sync/atomic"
)

const ()

type Arena struct {
	cap  int64
	cur  uint32 //current offset
	growby uint32
	buf  []byte //buffer
}

func newArena(cap int64, growby) *Arena {
	out := &Arena{
		cap: cap,
		cur:   1,
		buf: make([]byte, cap),
		growby: 1<<30,
	}
	return out
}

func (s *Arena) allocate(size uint32) (offset uint32) {
	//offset = s.cur + size
	offset = atomic.AddUint32(&s.cur, size)
	if offset <= uint32(s.cap) {
		return offset - size
	}

	//enlarge
	togrow:=s.growby
	if size > s.growby{
		togrow = size;
	} 
	newBuf = make([]byte, cap+int(togrow))
	AssertTrue(len(s.buf) == copy(newBuf, s.buf))
	s.buf = newBuf
	s.cap = len(s.buf)
	return offset - size

}

func (s *Arena) getsize()uint32{
	return atomic.LoadUint32(&s.cur)
}

func (s *Arena) getcapacity()int64{
	return atomic.LoadUint32(&s.cap)
}

func AssertTrue(b bool) {
	if !b{
		log.Fatalf("%+v", errors.Errorf("Assert failed"))
	}
}