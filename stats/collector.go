package stats

import "fmt"

// Collector collects and stores stats of RPSP game
type Collector struct {
	WinsCounter   *Counter
	LossesCounter *Counter
	DrawsCounter  *Counter
}

// NewCollector returns new collector
func NewCollector() *Collector {
	collector := &Collector{
		WinsCounter:   NewCounter(),
		LossesCounter: NewCounter(),
		DrawsCounter:  NewCounter(),
	}
	return collector
}

// PrintStats prints a nice string representation of stats
func (c *Collector) PrintStats() string {
	return fmt.Sprintf("W%d L%d D%d",
		c.WinsCounter.Count,
		c.LossesCounter.Count,
		c.DrawsCounter.Count)
}
