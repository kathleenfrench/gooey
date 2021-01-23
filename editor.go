package gooey

import (
	"github.com/AlecAivazis/survey/v2"
)

// Editor represents possible editor options in gooey
type Editor string

const (
	// VimEditor represents the vim editor
	VimEditor Editor = "vim"
	// SublimeEditor represents the sublime editor
	SublimeEditor Editor = "subl"
	// AtomEditor represents the sublime editor
	AtomEditor Editor = "atom"
	// NanoEditor represents the nano editor
	NanoEditor Editor = "nano"
)

// EditorPromptBuilder wraps inputs for gooey text editor prompts
type EditorPromptBuilder struct {
	Label        string
	DefaultValue string
	Editor       Editor
	HelpText     string
}

// String converts the editor to a string type
func (e Editor) String() string {
	return string(e)
}

// NewTextEditorInputPrompt creates a new editor prompt builder
func NewTextEditorInputPrompt(label string, defaultText string, editor Editor) (*EditorPromptBuilder, error) {
	tp := &EditorPromptBuilder{
		Label:        label,
		DefaultValue: defaultText,
		Editor:       editor,
	}

	err := tp.Valid()
	if err != nil {
		return nil, err
	}

	return tp, nil
}

// Valid checks whether the text editor prompt object has the minimum required values set to run
func (tp *EditorPromptBuilder) Valid() error {
	switch {
	case tp.Label == "":
		return genMissingFieldError("label")
	case !tp.Editor.Valid():
		return invalidFieldValueError("editor")
	default:
		return nil
	}
}

// Valid checks for a valid edditor
func (e Editor) Valid() bool {
	switch e {
	case VimEditor, NanoEditor, AtomEditor, SublimeEditor:
		return true
	default:
		return false
	}
}

// LaunchWithResponse launches a temporary file with a text editor, captures the input on save, and removes the tmp file while closing the editor
func (tp *EditorPromptBuilder) LaunchWithResponse() (content string) {
	prompt := &survey.Editor{
		Message:       tp.Label,
		Default:       tp.DefaultValue,
		AppendDefault: true,
		HideDefault:   true,
		Editor:        tp.Editor.String(),
		Help:          tp.HelpText,
	}

	survey.AskOne(prompt, &content)
	return content
}
