package yml

import (
	"code.lstaas.com/lightspeed/atom"
	"code.lstaas.com/lightspeed/atom/app"
	"github.com/MrVegeta/go-playground/analyst/infra/conf/serial"
)

func init() {
	atom.Must(app.RegisterConfigLoader(&app.ConfigFormat{
		Name:      "yml",
		Extension: []string{"yml", "yaml"},
		Loader:    serial.LoadCoreYmlConfig,
	}))
}
