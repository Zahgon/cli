package cli

import (
	"time"
)

type Args interface {
	Get(n int) string

	First() string

	Tail() []string

	Len() int

	Present() bool

	Slice() []string
}

type stringSliceArgs struct {
	v []string
}

func (a *stringSliceArgs) Get(n int) string { _ = "STUB: not implemented"; return "" }

func (a *stringSliceArgs) First() string { _ = "STUB: not implemented"; return "" }

func (a *stringSliceArgs) Tail() []string { _ = "STUB: not implemented"; return nil }

func (a *stringSliceArgs) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *stringSliceArgs) Present() bool { _ = "STUB: not implemented"; return false }

func (a *stringSliceArgs) Slice() []string { _ = "STUB: not implemented"; return nil }

type Argument interface {
	HasName(string) bool

	Parse([]string) ([]string, error)

	Usage() string

	Get() any
}

var AnyArguments = []Argument{
	&StringArgs{
		Max: -1,
	},
}

type ArgumentBase[T any, C any, VC ValueCreator[T, C]] struct {
	Name        string `json:"name"`
	Value       T      `json:"value"`
	Destination *T     `json:"-"`
	UsageText   string `json:"usageText"`
	Config      C      `json:"config"`

	value *T
}

func (a *ArgumentBase[T, C, VC]) HasName(s string) bool { _ = "STUB: not implemented"; return false }

func (a *ArgumentBase[T, C, VC]) Usage() string { _ = "STUB: not implemented"; return "" }

func (a *ArgumentBase[T, C, VC]) Parse(s []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ArgumentBase[T, C, VC]) Get() any { _ = "STUB: not implemented"; return *new(any) }

type ArgumentsBase[T any, C any, VC ValueCreator[T, C]] struct {
	Name        string `json:"name"`
	Value       T      `json:"value"`
	Destination *[]T   `json:"-"`
	UsageText   string `json:"usageText"`
	Min         int    `json:"minTimes"`
	Max         int    `json:"maxTimes"`
	Config      C      `json:"config"`

	values []T
}

func (a *ArgumentsBase[T, C, VC]) HasName(s string) bool { _ = "STUB: not implemented"; return false }

func (a *ArgumentsBase[T, C, VC]) Usage() string { _ = "STUB: not implemented"; return "" }

func (a *ArgumentsBase[T, C, VC]) Parse(s []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *ArgumentsBase[T, C, VC]) Get() any { _ = "STUB: not implemented"; return *new(any) }

type (
	FloatArg      = ArgumentBase[float64, NoConfig, floatValue[float64]]
	Float32Arg    = ArgumentBase[float32, NoConfig, floatValue[float32]]
	Float64Arg    = ArgumentBase[float64, NoConfig, floatValue[float64]]
	IntArg        = ArgumentBase[int, IntegerConfig, intValue[int]]
	Int8Arg       = ArgumentBase[int8, IntegerConfig, intValue[int8]]
	Int16Arg      = ArgumentBase[int16, IntegerConfig, intValue[int16]]
	Int32Arg      = ArgumentBase[int32, IntegerConfig, intValue[int32]]
	Int64Arg      = ArgumentBase[int64, IntegerConfig, intValue[int64]]
	StringArg     = ArgumentBase[string, StringConfig, stringValue]
	StringMapArgs = ArgumentBase[map[string]string, StringConfig, StringMap]
	TimestampArg  = ArgumentBase[time.Time, TimestampConfig, timestampValue]
	UintArg       = ArgumentBase[uint, IntegerConfig, uintValue[uint]]
	Uint8Arg      = ArgumentBase[uint8, IntegerConfig, uintValue[uint8]]
	Uint16Arg     = ArgumentBase[uint16, IntegerConfig, uintValue[uint16]]
	Uint32Arg     = ArgumentBase[uint32, IntegerConfig, uintValue[uint32]]
	Uint64Arg     = ArgumentBase[uint64, IntegerConfig, uintValue[uint64]]

	FloatArgs     = ArgumentsBase[float64, NoConfig, floatValue[float64]]
	Float32Args   = ArgumentsBase[float32, NoConfig, floatValue[float32]]
	Float64Args   = ArgumentsBase[float64, NoConfig, floatValue[float64]]
	IntArgs       = ArgumentsBase[int, IntegerConfig, intValue[int]]
	Int8Args      = ArgumentsBase[int8, IntegerConfig, intValue[int8]]
	Int16Args     = ArgumentsBase[int16, IntegerConfig, intValue[int16]]
	Int32Args     = ArgumentsBase[int32, IntegerConfig, intValue[int32]]
	Int64Args     = ArgumentsBase[int64, IntegerConfig, intValue[int64]]
	StringArgs    = ArgumentsBase[string, StringConfig, stringValue]
	TimestampArgs = ArgumentsBase[time.Time, TimestampConfig, timestampValue]
	UintArgs      = ArgumentsBase[uint, IntegerConfig, uintValue[uint]]
	Uint8Args     = ArgumentsBase[uint8, IntegerConfig, uintValue[uint8]]
	Uint16Args    = ArgumentsBase[uint16, IntegerConfig, uintValue[uint16]]
	Uint32Args    = ArgumentsBase[uint32, IntegerConfig, uintValue[uint32]]
	Uint64Args    = ArgumentsBase[uint64, IntegerConfig, uintValue[uint64]]
)

func (c *Command) getArgValue(name string) any { _ = "STUB: not implemented"; return *new(any) }

func arg[T any](name string, c *Command) T { _ = "STUB: not implemented"; return *new(T) }

func (c *Command) StringArg(name string) string { _ = "STUB: not implemented"; return "" }

func (c *Command) StringArgs(name string) []string { _ = "STUB: not implemented"; return nil }

func (c *Command) FloatArg(name string) float64 { _ = "STUB: not implemented"; return 0 }

func (c *Command) FloatArgs(name string) []float64 { _ = "STUB: not implemented"; return nil }

func (c *Command) Float32Arg(name string) float32 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Float32Args(name string) []float32 { _ = "STUB: not implemented"; return nil }

func (c *Command) Float64Arg(name string) float64 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Float64Args(name string) []float64 { _ = "STUB: not implemented"; return nil }

func (c *Command) IntArg(name string) int { _ = "STUB: not implemented"; return 0 }

func (c *Command) IntArgs(name string) []int { _ = "STUB: not implemented"; return nil }

func (c *Command) Int8Arg(name string) int8 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Int8Args(name string) []int8 { _ = "STUB: not implemented"; return nil }

func (c *Command) Int16Arg(name string) int16 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Int16Args(name string) []int16 { _ = "STUB: not implemented"; return nil }

func (c *Command) Int32Arg(name string) int32 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Int32Args(name string) []int32 { _ = "STUB: not implemented"; return nil }

func (c *Command) Int64Arg(name string) int64 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Int64Args(name string) []int64 { _ = "STUB: not implemented"; return nil }

func (c *Command) UintArg(name string) uint { _ = "STUB: not implemented"; return 0 }

func (c *Command) Uint8Arg(name string) uint8 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Uint16Arg(name string) uint16 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Uint32Arg(name string) uint32 { _ = "STUB: not implemented"; return 0 }

func (c *Command) Uint64Arg(name string) uint64 { _ = "STUB: not implemented"; return 0 }

func (c *Command) UintArgs(name string) []uint { _ = "STUB: not implemented"; return nil }

func (c *Command) Uint8Args(name string) []uint8 { _ = "STUB: not implemented"; return nil }

func (c *Command) Uint16Args(name string) []uint16 { _ = "STUB: not implemented"; return nil }

func (c *Command) Uint32Args(name string) []uint32 { _ = "STUB: not implemented"; return nil }

func (c *Command) Uint64Args(name string) []uint64 { _ = "STUB: not implemented"; return nil }

func (c *Command) TimestampArg(name string) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *Command) TimestampArgs(name string) []time.Time { _ = "STUB: not implemented"; return nil }
