package sync

import "sync"

type Counter struct {
	mutex sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}
