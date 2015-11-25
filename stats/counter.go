package stats

// Counter contains counter and channel
type Counter struct {
	Count   int
	Channel chan int
}

// NewCounter return a Counter
func NewCounter() *Counter {
	counter := &Counter{
		Count:   0,
		Channel: make(chan int),
	}
	go counter.statAgregator()
	return counter
}

func (c *Counter) statAgregator() {
	for {
		increment := <-c.Channel
		c.Count = c.Count + increment
	}
}
