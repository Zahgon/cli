package cli

type (
	IntSlice       = SliceBase[int, IntegerConfig, intValue[int]]
	Int8Slice      = SliceBase[int8, IntegerConfig, intValue[int8]]
	Int16Slice     = SliceBase[int16, IntegerConfig, intValue[int16]]
	Int32Slice     = SliceBase[int32, IntegerConfig, intValue[int32]]
	Int64Slice     = SliceBase[int64, IntegerConfig, intValue[int64]]
	IntSliceFlag   = FlagBase[[]int, IntegerConfig, IntSlice]
	Int8SliceFlag  = FlagBase[[]int8, IntegerConfig, Int8Slice]
	Int16SliceFlag = FlagBase[[]int16, IntegerConfig, Int16Slice]
	Int32SliceFlag = FlagBase[[]int32, IntegerConfig, Int32Slice]
	Int64SliceFlag = FlagBase[[]int64, IntegerConfig, Int64Slice]
)

var (
	NewIntSlice   = NewSliceBase[int, IntegerConfig, intValue[int]]
	NewInt8Slice  = NewSliceBase[int8, IntegerConfig, intValue[int8]]
	NewInt16Slice = NewSliceBase[int16, IntegerConfig, intValue[int16]]
	NewInt32Slice = NewSliceBase[int32, IntegerConfig, intValue[int32]]
	NewInt64Slice = NewSliceBase[int64, IntegerConfig, intValue[int64]]
)

func (cmd *Command) IntSlice(name string) []int { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Int8Slice(name string) []int8 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Int16Slice(name string) []int16 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Int32Slice(name string) []int32 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Int64Slice(name string) []int64 { _ = "STUB: not implemented"; return nil }
