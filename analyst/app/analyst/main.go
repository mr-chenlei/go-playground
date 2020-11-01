package main

import (
	"flag"
	"fmt"
	"os"

	"code.lstaas.com/lightspeed/atom/app"
	"code.lstaas.com/lightspeed/atom/log/errors"
	"github.com/MrVegeta/go-playground/analyst/common"
	"github.com/MrVegeta/go-playground/analyst/core"

	_ "github.com/MrVegeta/go-playground/analyst/app/analyst/distro/all"
)

var (
	config  = flag.String("config", "analyst.yml", "Config file for worker.")
	format  = flag.String("format", "yml", "Format of input file.")
	version = flag.Bool("version", false, "Show current version")
	build   = flag.Bool("build", false, "Show current build")
)

func main() {
	flag.Parse()

	if *version {
		common.PrintVersion()
	}
	if *build {
		common.PrintBuild()
	}
	if *version || *build {
		return
	}

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
