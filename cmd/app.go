package cmd

import (
	"github.com/mneumi/snt/internal/commands"
	"github.com/mneumi/snt/internal/common"
	"github.com/urfave/cli/v2"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:                 common.AppName,
		UsageText:            common.AppDesc,
		Version:              common.AppVersion,
		EnableBashCompletion: true,
	}

	app.Commands = []*cli.Command{
		commands.DoctorCommand,
		commands.UpgradeCommand,
		commands.FixCommand,
	}
}
