package process

import (
	"sync"
)

type Counter struct {
	sync.Mutex
	val int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.val++
}

func (c *Counter) Dec() {
	c.Lock()
	defer c.Unlock()
	c.val--
}

func (c *Counter) Val() int64 {
	c.Lock()
	defer c.Unlock()
	return c.val
}

func (c *Counter) Float() float64 {
	c.Lock()
	defer c.Unlock()
	return float64(c.val)
}
