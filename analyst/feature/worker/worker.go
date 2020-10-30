package worker

import (
	"context"

	"code.lstaas.com/lightspeed/atom/features"
)

// Handler provides interfaces for testing network performance.
type Handler interface {
	features.Feature
	Execute(ctx context.Context) error
}

// HandlerType returns the type of Handler interface. Can be used to implement atom.HasType
func HandlerType() interface{} {
	return (*Handler)(nil)
}
