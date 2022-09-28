package common

import (
	"container/ring"
	"log"
	"sync"
)

type RingBuffer struct {
	ring *ring.Ring
	mu   sync.RWMutex
}

func NewRingBuffer(size int) *RingBuffer {
	r := RingBuffer{
		ring: ring.New(size),
	}

	return &r
}

func (r *RingBuffer) Size() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.ring.Len()
}

func (r *RingBuffer) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, x := range p {
		r.ring.Value = x
		r.ring = r.ring.Next()
	}

	return len(p), nil
}

func (r *RingBuffer) Read(b []byte) (n int, err error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	i := 0

	r.ring.Do(func(p any) {
		if p == nil {
			return
		}

		x, ok := p.(byte)
		if !ok {
			log.Fatalf("failed to cast p=%#+v to byte", p)
		}

		b[i] = x

		i++
	})

	return i, err
}
