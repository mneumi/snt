package utils

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-version"
)

func CheckGoVersionGte(targetVer string) error {
	currentVer, err := extractGoVersion()
	if err != nil {
		return err
	}

	return checkGoVersionGte(currentVer, targetVer)
}

func extractGoVersion() (string, error) {
	ver, err := ExecuteCommand("go", "env", "GOVERSION")
	if err != nil {
		return "", err
	}

	// go1.20.6 => 1.20.6
	return strings.ReplaceAll(ver, "go", ""), nil
}

func checkGoVersionGte(currentVer string, targetVer string) error {
	v, err := version.NewVersion(currentVer)
	if err != nil {
		return err
	}

	condition := fmt.Sprintf(">= %s", targetVer)
	constraints, err := version.NewConstraint(condition)
	if err != nil {
		return err
	}

	if !constraints.Check(v) {
		return fmt.Errorf("go sdk version must be greater or equal than %s. current is %s", targetVer, currentVer)
	}

	return nil
}
