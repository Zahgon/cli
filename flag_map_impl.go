package cli

type MapBase[T any, C any, VC ValueCreator[T, C]] struct {
	dict             *map[string]T
	hasBeenSet       bool
	value            Value
	multiValueConfig multiValueParsingConfig
}

func (i MapBase[T, C, VC]) Create(val map[string]T, p *map[string]T, c C) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func NewMapBase[T any, C any, VC ValueCreator[T, C]](defaults map[string]T) *MapBase[T, C, VC] {
	_ = "STUB: not implemented"
	return nil
}

func (i *MapBase[T, C, VC]) setMultiValueParsingConfig(c multiValueParsingConfig) {
	_ = "STUB: not implemented"
	return
}

func (i *MapBase[T, C, VC]) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (i *MapBase[T, C, VC]) String() string { _ = "STUB: not implemented"; return "" }

func (i *MapBase[T, C, VC]) Serialize() string { _ = "STUB: not implemented"; return "" }

func (i *MapBase[T, C, VC]) Value() map[string]T { _ = "STUB: not implemented"; return nil }

func (i *MapBase[T, C, VC]) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (i MapBase[T, C, VC]) ToString(t map[string]T) string { _ = "STUB: not implemented"; return "" }

func sortedKeys[T any](dict map[string]T) []string { _ = "STUB: not implemented"; return nil }
