package gooey

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// InputPromptBuilder wraps inputs for gooey input prompts
type InputPromptBuilder struct {
	Label        string
	DefaultValue string
	HelpText     string
}

// InputPromptWithResponse accepts a user's typed input to a question as a response
func InputPromptWithResponse(label string, defaultVal string, disableClear bool) string {
	if !disableClear {
		ClearScreen()
	}

	var response string
	prompt := &survey.Input{
		Message: label,
		Default: defaultVal,
		Help:    ":q or :quit to exit",
	}

	survey.AskOne(prompt, &response)

	response = strings.TrimSpace(response)
	if wantsToExit(response) {
		os.Exit(0)
	}

	return response
}
