package cmd

import (
	"os"

	"github.com/mneumi/snt/internal/printer"
)

func Run() {
	if err := app.Run(os.Args); err != nil {
		printer.PrintError(err.Error())
		os.Exit(-1)
	}
}
