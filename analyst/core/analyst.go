package core

import (
	"code.lstaas.com/lightspeed/atom/app"
	"code.lstaas.com/lightspeed/atom/features"
	"code.lstaas.com/lightspeed/atom/log/errors"
	"code.lstaas.com/lightspeed/atom/serial"
)

// Instance combines all functionality in Magnetar.
type Instance struct {
	app.Instance
}

// New ...
func New(config *Config) (*Instance, error) {
	var server = &Instance{}
	for _, featureSettings := range config.Features {
		settings, err := serial.UnmarshalAny(featureSettings)
		if err != nil {
			return nil, err
		}
		obj, err := app.CreateObject(server, settings)
		if err != nil {
			return nil, err
		}
		if feature, ok := obj.(features.Feature); ok {
			if err := server.AddFeature(feature); err != nil {
				return nil, err
			}
		}
	}

	if server.Unresolved() {
		return nil, errors.New("not all dependency are resolved")
	}

	return server, nil
}

// Start overrides app.Instance.Start.
func (s *Instance) Start() error {
	err := s.Instance.Start()
	if err != nil {
		return err
	}
	errors.New("analyst started", "version", "").AtInfo().WriteToLog()
	return nil
}
