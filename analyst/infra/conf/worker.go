package conf

import "github.com/MrVegeta/go-playground/analyst/core/app/worker"

// WorkerConfig ...
type WorkerConfig struct {
	Path     *string `yaml:"path"`
	Template *string `yaml:"template"`
}

// Build ...
func (c *WorkerConfig) Build() (*worker.Config, error) {
	cfg := &worker.Config{}
	if c.Path != nil {
		cfg.Path = *c.Path
	}
	if c.Template != nil {
		cfg.Template = *c.Template
	}
	return cfg, nil
}
