package sync

import "sync"

type Counter struct {
	sync  sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.sync.Lock()
	defer c.sync.Unlock()
	c.value++
}
