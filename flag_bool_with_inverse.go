package cli

import (
	"context"
)

var DefaultInverseBoolPrefix = "no-"

type BoolWithInverseFlag struct {
	Name             string                                      `json:"name"`
	Category         string                                      `json:"category"`
	DefaultText      string                                      `json:"defaultText"`
	HideDefault      bool                                        `json:"hideDefault"`
	Usage            string                                      `json:"usage"`
	Sources          ValueSourceChain                            `json:"-"`
	Required         bool                                        `json:"required"`
	Hidden           bool                                        `json:"hidden"`
	Local            bool                                        `json:"local"`
	Value            bool                                        `json:"defaultValue"`
	Destination      *bool                                       `json:"-"`
	Aliases          []string                                    `json:"aliases"`
	TakesFile        bool                                        `json:"takesFileArg"`
	Action           func(context.Context, *Command, bool) error `json:"-"`
	OnlyOnce         bool                                        `json:"onlyOnce"`
	Validator        func(bool) error                            `json:"-"`
	ValidateDefaults bool                                        `json:"validateDefaults"`
	Config           BoolConfig                                  `json:"config"`
	InversePrefix    string                                      `json:"invPrefix"`

	count      int
	hasBeenSet bool
	applied    bool
	value      Value
	pset       bool
	nset       bool
}

func (bif *BoolWithInverseFlag) IsSet() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (bif *BoolWithInverseFlag) RunAction(ctx context.Context, cmd *Command) error {
	_ = "STUB: not implemented"
	return nil
}

func (bif *BoolWithInverseFlag) IsLocal() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) inversePrefix() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) PreParse() error { _ = "STUB: not implemented"; return nil }

func (bif *BoolWithInverseFlag) PostParse() error { _ = "STUB: not implemented"; return nil }

func (bif *BoolWithInverseFlag) Set(name, val string) error { _ = "STUB: not implemented"; return nil }

func (bif *BoolWithInverseFlag) Names() []string { _ = "STUB: not implemented"; return nil }

func (bif *BoolWithInverseFlag) IsRequired() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) IsVisible() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) String() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) Count() int { _ = "STUB: not implemented"; return 0 }

func (bif *BoolWithInverseFlag) GetDefaultText() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) GetCategory() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) SetCategory(c string) { _ = "STUB: not implemented"; return }

func (bif *BoolWithInverseFlag) GetUsage() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) GetEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (bif *BoolWithInverseFlag) GetValue() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) TakesValue() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) IsDefaultVisible() bool { _ = "STUB: not implemented"; return false }

func (bif *BoolWithInverseFlag) TypeName() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) SchemaType() string { _ = "STUB: not implemented"; return "" }

func (bif *BoolWithInverseFlag) SchemaItemsType() string { _ = "STUB: not implemented"; return "" }
