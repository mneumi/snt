package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecuteCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout

	err := cmd.Run()
	output := strings.TrimSpace(stdout.String())

	return output, err
}
