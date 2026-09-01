package cli

import "context"

type ShellCompleteFunc func(context.Context, *Command)

type BeforeFunc func(context.Context, *Command) (context.Context, error)

type AfterFunc func(context.Context, *Command) error

type ActionFunc func(context.Context, *Command) error

type ArgValidatorFunc func(context.Context, *Command) error

type CommandNotFoundFunc func(context.Context, *Command, string)

type ConfigureShellCompletionCommand func(*Command)

type OnUsageErrorFunc func(ctx context.Context, cmd *Command, err error, isSubcommand bool) error

type InvalidFlagAccessFunc func(context.Context, *Command, string)

type ExitErrHandlerFunc func(context.Context, *Command, error)

type FlagStringFunc func(Flag) string

type FlagNamePrefixFunc func(fullName []string, placeholder string) string

type FlagEnvHintFunc func(envVars []string, str string) string

type FlagFileHintFunc func(filePath, str string) string
