package cli

import (
	"context"
	"flag"
)

type Value interface {
	flag.Value
	flag.Getter
}

type boolFlag interface {
	IsBoolFlag() bool
}

type multiValueParsingConfig struct {
	SliceFlagSeparator string

	DisableSliceFlagSeparator bool

	MapFlagKeyValueSeparator string
}

type multiValueParsingConfigSetter interface {
	setMultiValueParsingConfig(c multiValueParsingConfig)
}

type ValueCreator[T any, C any] interface {
	Create(T, *T, C) Value
	ToString(T) string
}

type NoConfig struct{}

type FlagBase[T any, C any, VC ValueCreator[T, C]] struct {
	Name             string                                   `json:"name"`
	Category         string                                   `json:"category"`
	DefaultText      string                                   `json:"defaultText"`
	HideDefault      bool                                     `json:"hideDefault"`
	Usage            string                                   `json:"usage"`
	Sources          ValueSourceChain                         `json:"-"`
	Required         bool                                     `json:"required"`
	Hidden           bool                                     `json:"hidden"`
	Local            bool                                     `json:"local"`
	Value            T                                        `json:"defaultValue"`
	Destination      *T                                       `json:"-"`
	Aliases          []string                                 `json:"aliases"`
	TakesFile        bool                                     `json:"takesFileArg"`
	Action           func(context.Context, *Command, T) error `json:"-"`
	Config           C                                        `json:"config"`
	OnlyOnce         bool                                     `json:"onlyOnce"`
	Validator        func(T) error                            `json:"-"`
	ValidateDefaults bool                                     `json:"validateDefaults"`

	count      int
	hasBeenSet bool
	applied    bool
	creator    VC
	value      Value
}

func (f *FlagBase[T, C, V]) GetValue() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) TypeName() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) PostParse() error { _ = "STUB: not implemented"; return nil }

func (f *FlagBase[T, C, V]) setMultiValueParsingConfig(c multiValueParsingConfig) {
	_ = "STUB: not implemented"
	return
}

func (f *FlagBase[T, C, V]) PreParse() error { _ = "STUB: not implemented"; return nil }

func (f *FlagBase[T, C, V]) Set(_ string, val string) error { _ = "STUB: not implemented"; return nil }

func (f *FlagBase[T, C, V]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (f *FlagBase[T, C, V]) IsDefaultVisible() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, V]) String() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) IsSet() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, V]) Names() []string { _ = "STUB: not implemented"; return nil }

func (f *FlagBase[T, C, V]) IsRequired() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, V]) IsVisible() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, V]) GetCategory() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) SetCategory(c string) { _ = "STUB: not implemented"; return }

func (f *FlagBase[T, C, V]) GetUsage() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (f *FlagBase[T, C, V]) TakesValue() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, V]) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) RunAction(ctx context.Context, cmd *Command) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FlagBase[T, C, V]) SchemaType() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, V]) SchemaItemsType() string { _ = "STUB: not implemented"; return "" }

func (f *FlagBase[T, C, VC]) IsMultiValueFlag() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, VC]) IsLocal() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, VC]) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func (f *FlagBase[T, C, VC]) Count() int { _ = "STUB: not implemented"; return 0 }
