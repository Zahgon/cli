package cli

type StringFlag = FlagBase[string, StringConfig, stringValue]

type StringConfig struct {
	TrimSpace bool
}

type stringValue struct {
	destination *string
	trimSpace   bool
}

func (s stringValue) Create(val string, p *string, c StringConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (s stringValue) ToString(val string) string { _ = "STUB: not implemented"; return "" }

func (s *stringValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (s *stringValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (s *stringValue) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) String(name string) string { _ = "STUB: not implemented"; return "" }
