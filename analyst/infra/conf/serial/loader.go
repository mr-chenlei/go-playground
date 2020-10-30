package serial

import (
	"io"

	"github.com/MrVegeta/go-playground/analyst/infra/conf"
	yaml "gopkg.in/yaml.v2"
)

// LoadCoreYmlConfig ...
func LoadCoreYmlConfig(reader io.Reader) (interface{}, error) {
	ymlConfig := &conf.AnalystConfig{}

	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(ymlConfig); err != nil {
		return nil, err
	}

	config, err := ymlConfig.Build()
	if err != nil {
		return nil, err
	}

	return config, nil
}
