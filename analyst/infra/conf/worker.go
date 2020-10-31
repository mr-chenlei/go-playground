package conf

import "github.com/MrVegeta/go-playground/analyst/core/app/worker"

// WorkerConfig ...
type WorkerConfig struct {
	SDKLogPath      *string `yaml:"sdk-log-path"`
	FileType        *string `yaml:"file-type"`
	KeywordTemplate *string `yaml:"keyword-template"`
}

// Build ...
func (c *WorkerConfig) Build() (*worker.Config, error) {
	cfg := &worker.Config{}
	if c.SDKLogPath != nil {
		cfg.SDKLogPath = *c.SDKLogPath
	}
	if c.KeywordTemplate != nil {
		cfg.KeywordTemplate = *c.KeywordTemplate
	}
	return cfg, nil
}
