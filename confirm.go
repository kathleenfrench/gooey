package gooey

import "github.com/AlecAivazis/survey/v2"

// ConfirmPromptBuilder wraps inputs for gooey confirmation prompts
type ConfirmPromptBuilder struct {
	Label        string
	DefaultValue bool
	DisableClear bool
	HelpMessage  string
}

// NewConfirmPrompt creates a confirm prompt
func NewConfirmPrompt(label string, defaultValue bool, disableClear bool, helpMessage string) (*ConfirmPromptBuilder, error) {
	cp := &ConfirmPromptBuilder{
		Label:        label,
		DefaultValue: defaultValue,
		DisableClear: disableClear,
		HelpMessage:  helpMessage,
	}

	err := cp.Valid()
	if err != nil {
		return nil, err
	}

	return cp, nil
}

// Valid checks whether the minimum values have been set in order to return a valid prompt
func (cp *ConfirmPromptBuilder) Valid() error {
	if cp.Label == "" {
		return genMissingFieldError("label")
	}

	return nil
}

// LaunchWithResponse prompts the user for a yes/no response to a question, records then returns their response
func (cp *ConfirmPromptBuilder) LaunchWithResponse() (response bool) {
	if !cp.DisableClear {
		ClearScreen()
	}

	prompt := &survey.Confirm{
		Message: cp.Label,
		Default: cp.DefaultValue,
		Help:    cp.HelpMessage,
	}

	survey.AskOne(prompt, &response)
	return response
}
