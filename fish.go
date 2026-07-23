package cli

import (
	"io"
	"strings"
)

func (cmd *Command) ToFishCompletion() (string, error) { _ = "STUB: not implemented"; return "", nil }

type fishCommandCompletionTemplate struct {
	Command     *Command
	Completions []string
	AllCommands []string
}

func (cmd *Command) writeFishCompletionTemplate(w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareFishCommands(binary string, parent *Command) []string {
	_ = "STUB: not implemented"
	return nil
}

func prepareFishFlags(binary string, owner *Command) []string {
	_ = "STUB: not implemented"
	return nil
}

func fishAddFileFlag(flag Flag, completion *strings.Builder) { _ = "STUB: not implemented"; return }

func fishSubcommandHelper(binary string, command *Command, siblings []*Command) string {
	_ = "STUB: not implemented"
	return ""
}

func fishFlagHelper(binary string, command *Command) string { _ = "STUB: not implemented"; return "" }

func commandAncestry(command *Command) string { _ = "STUB: not implemented"; return "" }

func escapeSingleQuotes(input string) string { _ = "STUB: not implemented"; return "" }

var fishSingleQuoteReplacer = strings.NewReplacer(`\`, `\\`, `'`, `\'`)
