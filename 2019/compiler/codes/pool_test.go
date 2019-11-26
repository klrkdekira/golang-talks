package main

import (
	"bytes"
	"sync"
	"testing"
)

//NORMAL OMIT
func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		buf.WriteString("Hello")
		_ = buf
		buf.Reset()
	}
}

//NORMAL_END OMIT

//POOL OMIT
func BenchmarkPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	for i := 0; i < b.N; i++ {
		buf := pool.Get().(*bytes.Buffer)
		buf.WriteString("Hello")
		_ = buf
		buf.Reset()
		pool.Put(buf)
	}
}

//POOL_END OMIT
