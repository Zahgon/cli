package cli

type SliceBase[T any, C any, VC ValueCreator[T, C]] struct {
	slice                 *[]T
	hasBeenSet            bool
	value                 Value
	sliceSeparator        string
	disableSliceSeparator bool
}

func (i SliceBase[T, C, VC]) Create(val []T, p *[]T, c C) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func NewSliceBase[T any, C any, VC ValueCreator[T, C]](defaults ...T) *SliceBase[T, C, VC] {
	_ = "STUB: not implemented"
	return nil
}

func (i *SliceBase[T, C, VC]) setMultiValueParsingConfig(c multiValueParsingConfig) {
	_ = "STUB: not implemented"
	return
}

func (i *SliceBase[T, C, VC]) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (i *SliceBase[T, C, VC]) String() string { _ = "STUB: not implemented"; return "" }

func (i *SliceBase[T, C, VC]) Serialize() string { _ = "STUB: not implemented"; return "" }

func (i *SliceBase[T, C, VC]) Value() []T { _ = "STUB: not implemented"; return nil }

func (i *SliceBase[T, C, VC]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (i SliceBase[T, C, VC]) ToString(t []T) string { _ = "STUB: not implemented"; return "" }
