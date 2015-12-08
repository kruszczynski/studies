package stats

import "fmt"

// Collector collects and stores stats of RPSP game
type Collector struct {
	WinsCount   int32
	LossesCount int32
	DrawsCount  int32
}

// NewCollector returns new collector
func NewCollector() *Collector {
	collector := &Collector{
		WinsCount:   0,
		LossesCount: 0,
		DrawsCount:  0,
	}
	return collector
}

// PrintStats prints a nice string representation of stats
func (c *Collector) PrintStats() string {
	return fmt.Sprintf("W%d L%d D%d",
		c.WinsCount,
		c.LossesCount,
		c.DrawsCount)
}
