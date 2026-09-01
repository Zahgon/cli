package cli

type (
	UintSlice       = SliceBase[uint, IntegerConfig, uintValue[uint]]
	Uint8Slice      = SliceBase[uint8, IntegerConfig, uintValue[uint8]]
	Uint16Slice     = SliceBase[uint16, IntegerConfig, uintValue[uint16]]
	Uint32Slice     = SliceBase[uint32, IntegerConfig, uintValue[uint32]]
	Uint64Slice     = SliceBase[uint64, IntegerConfig, uintValue[uint64]]
	UintSliceFlag   = FlagBase[[]uint, IntegerConfig, UintSlice]
	Uint8SliceFlag  = FlagBase[[]uint8, IntegerConfig, Uint8Slice]
	Uint16SliceFlag = FlagBase[[]uint16, IntegerConfig, Uint16Slice]
	Uint32SliceFlag = FlagBase[[]uint32, IntegerConfig, Uint32Slice]
	Uint64SliceFlag = FlagBase[[]uint64, IntegerConfig, Uint64Slice]
)

var (
	NewUintSlice   = NewSliceBase[uint, IntegerConfig, uintValue[uint]]
	NewUint8Slice  = NewSliceBase[uint8, IntegerConfig, uintValue[uint8]]
	NewUint16Slice = NewSliceBase[uint16, IntegerConfig, uintValue[uint16]]
	NewUint32Slice = NewSliceBase[uint32, IntegerConfig, uintValue[uint32]]
	NewUint64Slice = NewSliceBase[uint64, IntegerConfig, uintValue[uint64]]
)

func (cmd *Command) UintSlice(name string) []uint { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Uint8Slice(name string) []uint8 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Uint16Slice(name string) []uint16 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Uint32Slice(name string) []uint32 { _ = "STUB: not implemented"; return nil }

func (cmd *Command) Uint64Slice(name string) []uint64 { _ = "STUB: not implemented"; return nil }

func getUintSlice[T uint | uint8 | uint16 | uint32 | uint64](cmd *Command, name string) []T {
	_ = "STUB: not implemented"
	return nil
}
