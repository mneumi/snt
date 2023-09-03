package utils

import (
	"fmt"

	"github.com/mneumi/snt/internal/common"
	"github.com/mneumi/snt/internal/printer"
)

func FindDependency(dependency *common.Dependency) (bool, string, error) {
	dependencyPath, err := ExecuteCommand("which", dependency.Name)
	if err != nil || dependencyPath == "" {
		return false, "", fmt.Errorf("%s: not found in $PATH", dependency.Name)
	}

	return true, dependencyPath, nil
}

func FixDependency(dependency *common.Dependency) error {
	switch dependency.Fix.Type {
	case common.FixMethodHint:
		printer.PrintCommon(dependency.Fix.Method)
	case common.FixMethodGoInstall:
		if _, err := ExecuteCommand("go", "install", fmt.Sprintf("%s@%s", dependency.Name, dependency.Version)); err != nil {
			return err
		}
		printer.PrintOK(dependency.Name)
	default:
		return fmt.Errorf("invalid install type %s", dependency.Fix.Type)
	}

	return nil
}
