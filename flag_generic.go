package cli

type GenericFlag = FlagBase[Value, NoConfig, genericValue]

type genericValue struct {
	val Value
}

func (f genericValue) Create(val Value, p *Value, c NoConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (f genericValue) ToString(b Value) string { _ = "STUB: not implemented"; return "" }

func (f *genericValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (f *genericValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (f *genericValue) String() string { _ = "STUB: not implemented"; return "" }

func (f *genericValue) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) Generic(name string) Value { _ = "STUB: not implemented"; return *new(Value) }
