package conf

import (
	"code.lstaas.com/lightspeed/atom/infra/conf"
	"code.lstaas.com/lightspeed/atom/serial"
	"github.com/MrVegeta/go-playground/analyst/core"
)

type AnalystConfig struct {
	Log      *conf.LogConfig `yaml:"log"`
	Analysis *WorkerConfig   `yaml:"worker"`
}

// Build ...
func (c *AnalystConfig) Build() (*core.Config, error) {
	cfg := &core.Config{}

	if c.Log != nil {
		logConf, err := c.Log.Build()
		if err != nil {
			return nil, err
		}
		cfg.Features = append(cfg.Features, serial.MarshalAny(logConf))
	}

	if c.Analysis != nil {
		oc, err := c.Analysis.Build()
		if err != nil {
			return nil, err
		}
		cfg.Features = append(cfg.Features, serial.MarshalAny(oc))
	}

	return cfg, nil
}
