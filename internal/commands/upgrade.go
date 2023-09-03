package commands

import (
	"github.com/mneumi/snt/internal/common"
	"github.com/mneumi/snt/internal/printer"
	"github.com/mneumi/snt/internal/utils"
	"github.com/urfave/cli/v2"
)

var UpgradeCommand = &cli.Command{
	Name:      "upgrade",
	UsageText: "升级 snt",
	Action: func(ctx *cli.Context) error {
		return upgrade()
	},
}

func upgrade() error {
	if _, err := utils.ExecuteCommand("go", "install", common.AppGitRepo+"@latest"); err != nil {
		return err
	}

	printer.PrintOK("Upgrade Success!")

	return nil
}
