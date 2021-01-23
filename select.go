package gooey

import "github.com/AlecAivazis/survey/v2"

// SelectPromptBuilder wraps common inputs for gooey dropdown selection prompts
type SelectPromptBuilder struct {
	Label        string
	Options      []string
	DefaultValue interface{}
	DisableClear bool
	MaxPageSize  int
	HelpMessage  string
}

// Valid checks whether the minimum values have been set in order to return a valid prompt
func (pb *SelectPromptBuilder) Valid() error {
	switch {
	case pb.Label == "":
		return genMissingFieldError("label")
	case len(pb.Options) == 0:
		return genMissingFieldError("options")
	default:
		return nil
	}
}

func (pb *SelectPromptBuilder) getPageSize() int {
	switch len(pb.Options) > pb.MaxPageSize {
	case true:
		return pb.MaxPageSize
	default:
		return len(pb.Options)
	}
}

// NewSelectPrompter creates and validates a new select prompt object
func NewSelectPrompter(
	label string,
	options []string,
	disableClear bool,
	maxSize int,
	helpMessage string,
	defaultValue interface{},
) (*SelectPromptBuilder, error) {
	sp := &SelectPromptBuilder{
		Label:        label,
		Options:      options,
		DefaultValue: defaultValue,
		HelpMessage:  helpMessage,
		DisableClear: disableClear,
		MaxPageSize:  maxSize,
	}

	err := sp.Valid()
	if err != nil {
		return nil, err
	}

	return sp, nil
}

// LaunchWithResponse launches the dropdown selection prompt then records and returns the user's selection
func (pb *SelectPromptBuilder) LaunchWithResponse() (selection string) {
	if pb.DisableClear {
		ClearScreen()
	}

	prompt := &survey.Select{
		Message:  pb.Label,
		Options:  pb.Options,
		PageSize: pb.getPageSize(),
	}

	if pb.DefaultValue != nil {
		prompt.Default = pb.DefaultValue
	}

	if pb.HelpMessage != "" {
		prompt.Help = pb.HelpMessage
	}

	survey.AskOne(prompt, &selection)
	return selection
}
