package cli

type (
	FloatSlice       = SliceBase[float64, NoConfig, floatValue[float64]]
	Float32Slice     = SliceBase[float32, NoConfig, floatValue[float32]]
	Float64Slice     = SliceBase[float64, NoConfig, floatValue[float64]]
	FloatSliceFlag   = FlagBase[[]float64, NoConfig, FloatSlice]
	Float32SliceFlag = FlagBase[[]float32, NoConfig, Float32Slice]
	Float64SliceFlag = FlagBase[[]float64, NoConfig, Float64Slice]
)

var (
	NewFloatSlice   = NewSliceBase[float64, NoConfig, floatValue[float64]]
	NewFloat32Slice = NewSliceBase[float32, NoConfig, floatValue[float32]]
	NewFloat64Slice = NewSliceBase[float64, NoConfig, floatValue[float64]]
)

func (cmd *Command) FloatSlice(name string) []float64 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Float32Slice(name string) []float32 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Float64Slice(name string) []float64 { _ = "STUB: not implemented"; return nil }
