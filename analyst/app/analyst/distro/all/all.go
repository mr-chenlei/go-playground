package all

import (
	// Other optional features
	_ "code.lstaas.com/lightspeed/atom/app/log"

	// YAML config support.
	_ "github.com/MrVegeta/go-playground/analyst/app/analyst/yml"

	// Load config from file or http(s)
	_ "code.lstaas.com/lightspeed/atom/confloader/external"
)
