package cli

type BoolFlag = FlagBase[bool, BoolConfig, boolValue]

type BoolConfig struct {
	Count *int
}

type boolValue struct {
	destination *bool
	count       *int
}

func (cmd *Command) Bool(name string) bool { _ = "STUB: not implemented"; return false }

func (b boolValue) Create(val bool, p *bool, c BoolConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (b boolValue) ToString(value bool) string { _ = "STUB: not implemented"; return "" }

func (b *boolValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (b *boolValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (b *boolValue) String() string { _ = "STUB: not implemented"; return "" }

func (b *boolValue) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }
