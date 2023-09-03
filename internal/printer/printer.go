package printer

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintGreen(format string, args ...any) {
	fmt.Println(color.GreenString(format, args...))
}

func PrintRed(format string, args ...any) {
	fmt.Println(color.RedString(format, args...))
}

func PrintCyan(format string, args ...any) {
	fmt.Println(color.CyanString(format, args...))
}

func PrintOK(format string, args ...any) {
	format = "✔  " + format
	PrintGreen(format, args...)
}

func PrintError(format string, args ...any) {
	format = "❗ " + format
	PrintRed(format, args...)
}

func PrintCommon(format string, args ...any) {
	PrintCyan(format, args...)
}
