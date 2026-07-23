package cli

import (
	"context"
	"io"
)

const (
	ignoreFlagPrefix = "test."

	commandContextKey = contextKey("cli.context")
)

type contextKey string

type Command struct {
	Name string `json:"name"`

	Aliases []string `json:"aliases"`

	Usage string `json:"usage"`

	UsageText string `json:"usageText"`

	ArgsUsage string `json:"argsUsage"`

	Version string `json:"version"`

	Description string `json:"description"`

	DefaultCommand string `json:"defaultCommand"`

	Category string `json:"category"`

	Commands []*Command `json:"commands"`

	Flags []Flag `json:"flags"`

	HideHelp bool `json:"hideHelp"`

	HideHelpCommand bool `json:"hideHelpCommand"`

	HideVersion bool `json:"hideVersion"`

	EnableShellCompletion bool `json:"-"`

	ShellCompletionCommandName string `json:"-"`

	ShellComplete ShellCompleteFunc `json:"-"`

	ConfigureShellCompletionCommand ConfigureShellCompletionCommand `json:"-"`

	Before BeforeFunc `json:"-"`

	After AfterFunc `json:"-"`

	Action ActionFunc `json:"-"`

	CommandNotFound CommandNotFoundFunc `json:"-"`

	OnUsageError OnUsageErrorFunc `json:"-"`

	InvalidFlagAccessHandler InvalidFlagAccessFunc `json:"-"`

	Hidden bool `json:"hidden"`

	Authors []any `json:"authors"`

	Copyright string `json:"copyright"`

	Reader io.Reader `json:"-"`

	Writer io.Writer `json:"-"`

	ErrWriter io.Writer `json:"-"`

	ExitErrHandler ExitErrHandlerFunc `json:"-"`

	Metadata map[string]any `json:"metadata"`

	ExtraInfo func() map[string]string `json:"-"`

	CustomRootCommandHelpTemplate string `json:"-"`

	SliceFlagSeparator string `json:"sliceFlagSeparator"`

	DisableSliceFlagSeparator bool `json:"disableSliceFlagSeparator"`

	MapFlagKeyValueSeparator string `json:"mapFlagKeyValueSeparator"`

	UseShortOptionHandling bool `json:"useShortOptionHandling"`

	Suggest bool `json:"suggest"`

	AllowExtFlags bool `json:"allowExtFlags"`

	SkipFlagParsing bool `json:"skipFlagParsing"`

	CustomHelpTemplate string `json:"-"`

	PrefixMatchCommands bool `json:"prefixMatchCommands"`

	SuggestCommandFunc SuggestCommandFunc `json:"-"`

	MutuallyExclusiveFlags []MutuallyExclusiveFlags `json:"mutuallyExclusiveFlags"`

	Arguments []Argument `json:"arguments"`

	ReadArgsFromStdin bool `json:"readArgsFromStdin"`

	StopOnNthArg *int `json:"stopOnNthArg"`

	categories CommandCategories

	flagCategories FlagCategories

	appliedFlags []Flag

	setFlags map[Flag]struct{}

	parent *Command

	parsedArgs Args

	isInError bool

	didSetupDefaults bool

	shellCompletion bool

	globaHelpFlagAdded bool

	globaVersionFlagAdded bool

	versionFlag Flag

	isCompletionCommand bool

	builtInHelp bool
}

func (cmd *Command) Command(name string) *Command { _ = "STUB: not implemented"; return nil }

func (cmd *Command) checkHelp() bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) allFlags() []Flag { _ = "STUB: not implemented"; return nil }

func (cmd *Command) useShortOptionHandling() bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) suggestFlagFromError(err error, commandName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cmd *Command) Names() []string { _ = "STUB: not implemented"; return nil }

func (cmd *Command) HasName(name string) bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) VisibleCategories() []CommandCategory { _ = "STUB: not implemented"; return nil }

func (cmd *Command) VisibleCommands() []*Command { _ = "STUB: not implemented"; return nil }

func (cmd *Command) VisibleFlagCategories() []VisibleFlagCategory {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *Command) VisibleFlags() []Flag { _ = "STUB: not implemented"; return nil }

func (cmd *Command) appendFlag(fl Flag) { _ = "STUB: not implemented"; return }

func (cmd *Command) VisiblePersistentFlags() []Flag { _ = "STUB: not implemented"; return nil }

func (cmd *Command) appendCommand(aCmd *Command) { _ = "STUB: not implemented"; return }

func (cmd *Command) handleExitCoder(ctx context.Context, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *Command) argsWithDefaultCommand(oldArgs Args) Args {
	_ = "STUB: not implemented"
	return *new(Args)
}

func (cmd *Command) Root() *Command { _ = "STUB: not implemented"; return nil }

func (cmd *Command) set(fName string, f Flag, val string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *Command) lFlag(name string) Flag { _ = "STUB: not implemented"; return *new(Flag) }

func (cmd *Command) hasPersistentFlagOnAncestor(fl Flag) bool {
	_ = "STUB: not implemented"
	return false
}

func (cmd *Command) lookupFlag(name string) Flag { _ = "STUB: not implemented"; return *new(Flag) }

func (cmd *Command) lookupAppliedFlag(name string) Flag {
	_ = "STUB: not implemented"
	return *new(Flag)
}

func (cmd *Command) checkRequiredFlag(f Flag) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (cmd *Command) checkAllRequiredFlags() requiredFlagsErr {
	_ = "STUB: not implemented"
	return *new(requiredFlagsErr)
}

func (cmd *Command) checkRequiredFlags() requiredFlagsErr {
	_ = "STUB: not implemented"
	return *new(requiredFlagsErr)
}

func (cmd *Command) onInvalidFlag(ctx context.Context, name string) {
	_ = "STUB: not implemented"
	return
}

func (cmd *Command) NumFlags() int { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) setMultiValueParsingConfig(f Flag) { _ = "STUB: not implemented"; return }

func (cmd *Command) Set(name, value string) error { _ = "STUB: not implemented"; return nil }

func (cmd *Command) IsSet(name string) bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) LocalFlagNames() []string { _ = "STUB: not implemented"; return nil }

func (cmd *Command) FlagNames() []string { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Lineage() []*Command { _ = "STUB: not implemented"; return nil }

func (cmd *Command) FullName() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) Path() []string { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Walk(fn func(*Command) error) error { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Count(name string) int { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Value(name string) any { _ = "STUB: not implemented"; return *new(any) }

func (cmd *Command) Args() Args { _ = "STUB: not implemented"; return *new(Args) }

func (cmd *Command) NArg() int { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) runFlagActions(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
