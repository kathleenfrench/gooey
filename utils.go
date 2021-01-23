package gooey

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

// maxPageSize represents the default maximum number of dropdown options to display in gooey
const maxPageSize = 20

// ClearScreen clears the terminal UI
func ClearScreen() {
	print("\033[H\033[2J")
}

func wantsToExit(userInput string, exitKeys []string) bool {
	for _, k := range exitKeys {
		if userInput == k {
			return true
		}
	}

	return false
}

// Spin is a spinner used to indicate a pending process
var Spin = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

// Exit says by and exits the program
func Exit() {
	color.HiCyan("\nbye%s", emoji.Sprint(":wave:"))
	os.Exit(0)
}
