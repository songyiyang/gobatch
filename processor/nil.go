package processor

import (
	"context"
	"time"

	"github.com/songyiyang/gobatch/batch"
)

// Nil is a Processor that discards all data after a specified duration.
// It can be used as a mock Processor.
type Nil struct {
	Duration time.Duration
}

// Process discards all data sent to it after a certain amount of time.
func (p *Nil) Process(ctx context.Context, ps *batch.PipelineStage) {
	time.Sleep(p.Duration)
	ps.Close()
}
