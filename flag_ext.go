package cli

import (
	"flag"
)

type extFlag struct {
	f *flag.Flag
}

func (e *extFlag) PreParse() error { _ = "STUB: not implemented"; return nil }

func (e *extFlag) PostParse() error { _ = "STUB: not implemented"; return nil }

func (e *extFlag) Set(_ string, val string) error { _ = "STUB: not implemented"; return nil }

func (e *extFlag) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (e *extFlag) Names() []string { _ = "STUB: not implemented"; return nil }

func (e *extFlag) IsSet() bool { _ = "STUB: not implemented"; return false }

func (e *extFlag) String() string { _ = "STUB: not implemented"; return "" }

func (e *extFlag) IsVisible() bool { _ = "STUB: not implemented"; return false }

func (e *extFlag) TakesValue() bool { _ = "STUB: not implemented"; return false }

func (e *extFlag) GetUsage() string { _ = "STUB: not implemented"; return "" }

func (e *extFlag) GetValue() string { _ = "STUB: not implemented"; return "" }

func (e *extFlag) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

func (e *extFlag) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (e *extFlag) SchemaType() string { _ = "STUB: not implemented"; return "" }

func (e *extFlag) SchemaItemsType() string { _ = "STUB: not implemented"; return "" }
