package cli

import (
	"context"
	"io"
)

const (
	helpName  = "help"
	helpAlias = "h"
)

type HelpPrinterFunc func(w io.Writer, templ string, data any)

type HelpPrinterCustomFunc func(w io.Writer, templ string, data any, customFunc map[string]any)

var HelpPrinter HelpPrinterFunc = DefaultPrintHelp

var HelpPrinterCustom HelpPrinterCustomFunc = DefaultPrintHelpCustom

var VersionPrinter = DefaultPrintVersion

var ShowRootCommandHelp = DefaultShowRootCommandHelp

var ShowAppHelp = ShowRootCommandHelp

var ShowCommandHelp = DefaultShowCommandHelp

var ShowSubcommandHelp = DefaultShowSubcommandHelp

var UsageCommandHelp = "Shows a list of commands or help for one command"

var ArgsUsageCommandHelp = "[command]"

func buildHelpCommand(withAction bool) *Command { _ = "STUB: not implemented"; return nil }

func helpCommandAction(ctx context.Context, cmd *Command) error {
	_ = "STUB: not implemented"
	return nil
}

func ShowRootCommandHelpAndExit(cmd *Command, exitCode int) { _ = "STUB: not implemented"; return }

var ShowAppHelpAndExit = ShowRootCommandHelpAndExit

func DefaultShowRootCommandHelp(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func DefaultRootCommandComplete(ctx context.Context, cmd *Command) {
	_ = "STUB: not implemented"
	return
}

var DefaultAppComplete = DefaultRootCommandComplete

func printCommandSuggestions(commands []*Command, writer io.Writer) {
	_ = "STUB: not implemented"
	return
}

func cliArgContains(flagName string, args []string) bool { _ = "STUB: not implemented"; return false }

func printFlagSuggestions(lastArg string, flags []Flag, writer io.Writer) {
	_ = "STUB: not implemented"
	return
}

func DefaultCompleteWithFlags(ctx context.Context, cmd *Command) { _ = "STUB: not implemented"; return }

func ShowCommandHelpAndExit(ctx context.Context, cmd *Command, command string, code int) {
	_ = "STUB: not implemented"
	return
}

func DefaultShowCommandHelp(ctx context.Context, cmd *Command, commandName string) error {
	_ = "STUB: not implemented"
	return nil
}

func ShowSubcommandHelpAndExit(cmd *Command, exitCode int) { _ = "STUB: not implemented"; return }

func DefaultShowSubcommandHelp(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func ShowVersion(cmd *Command) { _ = "STUB: not implemented"; return }

func DefaultPrintVersion(cmd *Command) { _ = "STUB: not implemented"; return }

func handleTemplateError(err error) { _ = "STUB: not implemented"; return }

func DefaultPrintHelpCustom(out io.Writer, templ string, data any, customFuncs map[string]any) {
	_ = "STUB: not implemented"
	return
}

func DefaultPrintHelp(out io.Writer, templ string, data any) { _ = "STUB: not implemented"; return }

func checkVersion(cmd *Command) bool { _ = "STUB: not implemented"; return false }

func checkShellCompleteFlag(c *Command, arguments []string) (bool, []string) {
	_ = "STUB: not implemented"
	return false, nil
}

func shouldRunCompletion(cmd *Command) bool { _ = "STUB: not implemented"; return false }

func runCompletion(ctx context.Context, cmd *Command) { _ = "STUB: not implemented"; return }

func subtract(a, b int) int { _ = "STUB: not implemented"; return 0 }

func indent(spaces int, v string) string { _ = "STUB: not implemented"; return "" }

func nindent(spaces int, v string) string { _ = "STUB: not implemented"; return "" }

func wrap(input string, offset int, wrapAt int) string { _ = "STUB: not implemented"; return "" }

func wrapLine(input string, offset int, wrapAt int, padding string) string {
	_ = "STUB: not implemented"
	return ""
}

func offset(input string, fixed int) int { _ = "STUB: not implemented"; return 0 }

func offsetCommands(cmds []*Command, fixed int) int { _ = "STUB: not implemented"; return 0 }
