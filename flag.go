package cli

import (
	"context"
	"fmt"
	"time"
)

const defaultPlaceholder = "value"

const (
	defaultSliceFlagSeparator       = ","
	defaultMapFlagKeyValueSeparator = "="
	disableSliceFlagSeparator       = false
)

var slPfx = fmt.Sprintf("sl:::%d:::", time.Now().UTC().UnixNano())

var GenerateShellCompletionFlag Flag = &BoolFlag{
	Name:   "generate-shell-completion",
	Hidden: true,
}

var VersionFlag Flag = &BoolFlag{
	Name:        "version",
	Aliases:     []string{"v"},
	Usage:       "print the version",
	HideDefault: true,
	Local:       true,
}

var HelpFlag Flag = &BoolFlag{
	Name:        "help",
	Aliases:     []string{"h"},
	Usage:       "show help",
	HideDefault: true,
	Local:       true,
}

var FlagStringer FlagStringFunc = stringifyFlag

type Serializer interface {
	Serialize() string
}

var FlagNamePrefixer FlagNamePrefixFunc = prefixedNames

var FlagEnvHinter FlagEnvHintFunc = withEnvHint

var FlagFileHinter FlagFileHintFunc = withFileHint

type FlagsByName []Flag

func (f FlagsByName) Len() int { _ = "STUB: not implemented"; return 0 }

func (f FlagsByName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (f FlagsByName) Swap(i, j int) { _ = "STUB: not implemented"; return }

type ActionableFlag interface {
	RunAction(context.Context, *Command) error
}

type Flag interface {
	fmt.Stringer

	Get() any

	PreParse() error

	PostParse() error

	Set(string, string) error

	Names() []string

	IsSet() bool
}

type RequiredFlag interface {
	IsRequired() bool
}

type DocGenerationFlag interface {
	TakesValue() bool

	GetUsage() string

	GetValue() string

	GetDefaultText() string

	GetEnvVars() []string

	IsDefaultVisible() bool

	TypeName() string
}

type DocGenerationMultiValueFlag interface {
	DocGenerationFlag

	IsMultiValueFlag() bool
}

type SchemaTyper interface {
	SchemaType() string
}

type SchemaItemsTyper interface {
	SchemaItemsType() string
}

type Countable interface {
	Count() int
}

type VisibleFlag interface {
	IsVisible() bool
}

type CategorizableFlag interface {
	GetCategory() string

	SetCategory(string)
}

type LocalFlag interface {
	IsLocal() bool
}

func visibleFlags(fl []Flag) []Flag { _ = "STUB: not implemented"; return nil }

func FlagNames(name string, aliases []string) []string { _ = "STUB: not implemented"; return nil }

func hasFlag(flags []Flag, fl Flag) bool { _ = "STUB: not implemented"; return false }

func flagSplitMultiValues(val string, sliceSeparator string, disableSliceSeparator bool) []string {
	_ = "STUB: not implemented"
	return nil
}
