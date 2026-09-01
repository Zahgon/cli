package cli

type (
	IntFlag   = FlagBase[int, IntegerConfig, intValue[int]]
	Int8Flag  = FlagBase[int8, IntegerConfig, intValue[int8]]
	Int16Flag = FlagBase[int16, IntegerConfig, intValue[int16]]
	Int32Flag = FlagBase[int32, IntegerConfig, intValue[int32]]
	Int64Flag = FlagBase[int64, IntegerConfig, intValue[int64]]
)

type IntegerConfig struct {
	Base int
}

type intValue[T int | int8 | int16 | int32 | int64] struct {
	val  *T
	base int
}

func (i intValue[T]) Create(val T, p *T, c IntegerConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (i intValue[T]) ToString(b T) string { _ = "STUB: not implemented"; return "" }

func (i *intValue[T]) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (i *intValue[T]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (i *intValue[T]) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) Int(name string) int { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Int8(name string) int8 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Int16(name string) int16 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Int32(name string) int32 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Int64(name string) int64 { _ = "STUB: not implemented"; return 0 }

func getInt[T int | int8 | int16 | int32 | int64](cmd *Command, name string) T {
	_ = "STUB: not implemented"
	return *new(T)
}
