package commands

import (
	"github.com/mneumi/snt/internal/common"
	"github.com/mneumi/snt/internal/utils"
	"github.com/urfave/cli/v2"
)

var FixCommand = &cli.Command{
	Name:      "fix",
	UsageText: "环境依赖初始化",
	Action: func(ctx *cli.Context) error {
		if err := utils.CheckGoVersionGte(common.GoVersionRequire); err != nil {
			return err
		}

		if err := fixDependencies(common.Dependencies); err != nil {
			return err
		}

		return nil
	},
}

func fixDependencies(dependencies []*common.Dependency) error {
	for _, dependency := range dependencies {
		if exist, _, _ := utils.FindDependency(dependency); !exist {
			return utils.FixDependency(dependency)
		}
	}

	return nil
}
