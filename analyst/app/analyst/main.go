package main

import (
	"flag"
	"fmt"
	"os"

	"code.lstaas.com/lightspeed/atom/app"
	"code.lstaas.com/lightspeed/atom/log/errors"
	"github.com/MrVegeta/go-playground/analyst/core"

	_ "github.com/MrVegeta/go-playground/analyst/app/analyst/distro/all"
)

var (
	config = flag.String("config", "", "Config file for worker.")
	format = flag.String("format", "yml", "Format of input file.")
)

func main() {
	flag.Parse()

	server, err := startAnalyst()
	if err != nil {
		fmt.Println(err.Error())
		// Configuration error. Exit with a special value to prevent system from restarting.
		os.Exit(23)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
	defer server.Close()
}

func startAnalyst() (app.Server, error) {
	ci, err := app.LoadConfig(*format, *config)
	if err != nil {
		return nil, errors.NewF("failed to read config file: %s", *config).Base(err)
	}

	config, ok := ci.(*core.Config)
	if !ok {
		return nil, errors.New("invalid config")
	}

	server, err := core.New(config)
	if err != nil {
		return nil, errors.New("failed to create server").Base(err)
	}

	return server, nil
}
