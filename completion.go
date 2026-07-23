package cli

import (
	"embed"
	"fmt"
)

const (
	completionCommandName = "completion"

	completionFlag = "--generate-shell-completion"
)

type renderCompletion func(cmd *Command, appName string) (string, error)

var (
	//go:embed autocomplete
	autoCompleteFS embed.FS

	completionShells = []string{"bash", "zsh", "fish", "pwsh"}

	shellCompletions = map[string]renderCompletion{
		"bash": func(c *Command, appName string) (string, error) {
			b, err := autoCompleteFS.ReadFile("autocomplete/bash_autocomplete")
			return fmt.Sprintf(string(b), appName), err
		},
		"zsh": func(c *Command, appName string) (string, error) {
			b, err := autoCompleteFS.ReadFile("autocomplete/zsh_autocomplete")
			return fmt.Sprintf(string(b), appName), err
		},
		"fish": func(c *Command, appName string) (string, error) {
			b, err := autoCompleteFS.ReadFile("autocomplete/fish_autocomplete")
			return fmt.Sprintf(string(b), appName), err
		},
		"pwsh": func(c *Command, appName string) (string, error) {
			b, err := autoCompleteFS.ReadFile("autocomplete/powershell_autocomplete.ps1")
			return string(b), err
		},
	}
)

const completionDescription = `Output shell completion script for bash, zsh, fish, or Powershell.
Source the output to enable completion.

# .bashrc
source <($COMMAND completion bash)

# .zshrc
source <($COMMAND completion zsh)

# fish
$COMMAND completion fish > ~/.config/fish/completions/$COMMAND.fish

# Powershell
Output the script to path/to/autocomplete/$COMMAND.ps1 and run it.
`

func buildCompletionCommand(appName string) *Command { _ = "STUB: not implemented"; return nil }

func buildShellCompletionSubcommand(shell string, render renderCompletion, appName string) *Command {
	_ = "STUB: not implemented"
	return nil
}
