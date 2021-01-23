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
	DisableClear bool
	ExitKeys     []string
}

// NewInputPrompt creates and validates a new input prompt builder
func NewInputPrompt(label string, defaultValue string, disableClear bool, helpText string) (*InputPromptBuilder, error) {
	ip := &InputPromptBuilder{
		Label:        label,
		DefaultValue: defaultValue,
		HelpText:     helpText,
		DisableClear: disableClear,
	}

	err := ip.Valid()
	if err != nil {
		return nil, err
	}

	return ip, nil
}

// SetExitKeys allows for setting a slice of user inputs that will exit a terminal program
func (ip *InputPromptBuilder) SetExitKeys(keys []string) {
	ip.ExitKeys = keys
}

// Valid validates that the minimum values required to run the input prompt are present
func (ip *InputPromptBuilder) Valid() error {
	switch {
	case ip.Label == "":
		return genMissingFieldError("label")
	default:
		return nil
	}
}

// LaunchWithResponse launches the input prompt and accepts a user's string input to a question then returns the response
func (ip *InputPromptBuilder) LaunchWithResponse() (response string) {
	if ip.DisableClear {
		ClearScreen()
	}

	prompt := &survey.Input{
		Message: ip.Label,
		Default: ip.DefaultValue,
		Help:    ip.HelpText,
	}

	survey.AskOne(prompt, &response)
	response = strings.TrimSpace(response)

	if len(ip.ExitKeys) > 0 && wantsToExit(response, ip.ExitKeys) {
		os.Exit(0)
	}

	return response
}
