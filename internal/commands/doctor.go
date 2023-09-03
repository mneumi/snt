package commands

import (
	"github.com/mneumi/snt/internal/common"
	"github.com/mneumi/snt/internal/printer"
	"github.com/mneumi/snt/internal/utils"
	"github.com/urfave/cli/v2"
)

var DoctorCommand = &cli.Command{
	Name:      "doctor",
	UsageText: "环境依赖检查",
	Action: func(ctx *cli.Context) error {
		if err := checkGoVersionGte(); err != nil {
			return err
		}

		if err := checkDependencies(common.Dependencies); err != nil {
			return err
		}

		return nil
	},
}

func checkGoVersionGte() error {
	if err := utils.CheckGoVersionGte(common.GoVersionRequire); err != nil {
		return err
	}

	printer.PrintOK("Go SDK Version")

	return nil
}

func checkDependencies(dependencies []*common.Dependency) error {
	for _, dependency := range dependencies {
		if _, _, err := utils.FindDependency(dependency); err != nil {
			return err
		}

		printer.PrintOK(dependency.Name)
	}

	return nil
}
