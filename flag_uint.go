package cli

type (
	UintFlag   = FlagBase[uint, IntegerConfig, uintValue[uint]]
	Uint8Flag  = FlagBase[uint8, IntegerConfig, uintValue[uint8]]
	Uint16Flag = FlagBase[uint16, IntegerConfig, uintValue[uint16]]
	Uint32Flag = FlagBase[uint32, IntegerConfig, uintValue[uint32]]
	Uint64Flag = FlagBase[uint64, IntegerConfig, uintValue[uint64]]
)

type uintValue[T uint | uint8 | uint16 | uint32 | uint64] struct {
	val  *T
	base int
}

func (i uintValue[T]) Create(val T, p *T, c IntegerConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (i uintValue[T]) ToString(b T) string { _ = "STUB: not implemented"; return "" }

func (i *uintValue[T]) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (i *uintValue[T]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (i *uintValue[T]) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) Uint(name string) uint { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Uint8(name string) uint8 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Uint16(name string) uint16 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Uint32(name string) uint32 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Uint64(name string) uint64 { _ = "STUB: not implemented"; return 0 }

func getUint[T uint | uint8 | uint16 | uint32 | uint64](cmd *Command, name string) T {
	_ = "STUB: not implemented"
	return *new(T)
}
