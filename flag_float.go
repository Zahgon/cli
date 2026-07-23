package cli

type (
	FloatFlag   = FlagBase[float64, NoConfig, floatValue[float64]]
	Float32Flag = FlagBase[float32, NoConfig, floatValue[float32]]
	Float64Flag = FlagBase[float64, NoConfig, floatValue[float64]]
)

type floatValue[T float32 | float64] struct {
	val *T
}

func (f floatValue[T]) Create(val T, p *T, c NoConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (f floatValue[T]) ToString(b T) string { _ = "STUB: not implemented"; return "" }

func (f *floatValue[T]) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (f *floatValue[T]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (f *floatValue[T]) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) Float(name string) float64 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Float32(name string) float32 { _ = "STUB: not implemented"; return 0 }

func (cmd *Command) Float64(name string) float64 { _ = "STUB: not implemented"; return 0 }

func getFloat[T float32 | float64](cmd *Command, name string) T {
	_ = "STUB: not implemented"
	return *new(T)
}
