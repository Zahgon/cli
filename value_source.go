package cli

import (
	"fmt"
)

type ValueSource interface {
	fmt.Stringer
	fmt.GoStringer

	Lookup() (string, bool)
}

type EnvValueSource interface {
	IsFromEnv() bool
	Key() string
}

type MapSource interface {
	fmt.Stringer
	fmt.GoStringer

	Lookup(string) (any, bool)
}

type ValueSourceChain struct {
	Chain []ValueSource
}

func NewValueSourceChain(src ...ValueSource) ValueSourceChain {
	_ = "STUB: not implemented"
	return *new(ValueSourceChain)
}

func (vsc *ValueSourceChain) Append(other ValueSourceChain) { _ = "STUB: not implemented"; return }

func (vsc *ValueSourceChain) EnvKeys() []string { _ = "STUB: not implemented"; return nil }

func (vsc *ValueSourceChain) String() string { _ = "STUB: not implemented"; return "" }

func (vsc *ValueSourceChain) GoString() string { _ = "STUB: not implemented"; return "" }

func (vsc *ValueSourceChain) Lookup() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (vsc *ValueSourceChain) LookupWithSource() (string, ValueSource, bool) {
	_ = "STUB: not implemented"
	return "", *new(ValueSource), false
}

type envVarValueSource struct {
	key string
}

func (e *envVarValueSource) Lookup() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (e *envVarValueSource) IsFromEnv() bool { _ = "STUB: not implemented"; return false }

func (e *envVarValueSource) Key() string { _ = "STUB: not implemented"; return "" }

func (e *envVarValueSource) String() string   { _ = "STUB: not implemented"; return "" }
func (e *envVarValueSource) GoString() string { _ = "STUB: not implemented"; return "" }

func EnvVar(key string) ValueSource { _ = "STUB: not implemented"; return *new(ValueSource) }

func EnvVars(keys ...string) ValueSourceChain {
	_ = "STUB: not implemented"
	return *new(ValueSourceChain)
}

type fileValueSource struct {
	Path string
}

func (f *fileValueSource) Lookup() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (f *fileValueSource) String() string   { _ = "STUB: not implemented"; return "" }
func (f *fileValueSource) GoString() string { _ = "STUB: not implemented"; return "" }

func File(path string) ValueSource { _ = "STUB: not implemented"; return *new(ValueSource) }

func Files(paths ...string) ValueSourceChain {
	_ = "STUB: not implemented"
	return *new(ValueSourceChain)
}

type mapSource struct {
	name string
	m    map[any]any
}

func NewMapSource(name string, m map[any]any) MapSource {
	_ = "STUB: not implemented"
	return *new(MapSource)
}

func (ms *mapSource) String() string   { _ = "STUB: not implemented"; return "" }
func (ms *mapSource) GoString() string { _ = "STUB: not implemented"; return "" }

func (ms *mapSource) Lookup(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

type mapValueSource struct {
	key string
	ms  MapSource
}

func NewMapValueSource(key string, ms MapSource) ValueSource {
	_ = "STUB: not implemented"
	return *new(ValueSource)
}

func (mvs *mapValueSource) String() string { _ = "STUB: not implemented"; return "" }

func (mvs *mapValueSource) GoString() string { _ = "STUB: not implemented"; return "" }

func (mvs *mapValueSource) Lookup() (string, bool) { _ = "STUB: not implemented"; return "", false }
