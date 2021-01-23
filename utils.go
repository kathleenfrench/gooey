package gooey

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

func clearScreen() {
	print("\033[H\033[2J")
}

func wantsToExit(v string) bool {
	switch v {
	case ":q", ":quit":
		return true
	default:
		return false
	}
}

// Spin is a spinner used to indicate a pending process
var Spin = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

// Exit says by and exits the program
func Exit() {
	color.HiCyan("\nbye%s", emoji.Sprint(":wave:"))
	os.Exit(0)
}
