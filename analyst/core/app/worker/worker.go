package worker

import (
	"context"

	"code.lstaas.com/lightspeed/atom"
)

type Worker struct {
	Path     string
	Template string
}

func New(ctx context.Context, config *Config) (*Worker, error) {
	return &Worker{
		Path:     config.Path,
		Template: config.Template,
	}, nil
}
func init() {
	atom.Must(atom.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
