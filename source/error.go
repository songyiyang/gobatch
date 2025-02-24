package source

import (
	"context"

	"github.com/songyiyang/gobatch/batch"
)

// Error is a Source that returns an error and then closes immediately.
// It can be used as a mock Source.
type Error struct {
	Err error
}

// Read returns an error and then closes.
func (s *Error) Read(ctx context.Context, ps *batch.PipelineStage) {
	ps.Errors <- s.Err
	ps.Close()
}
