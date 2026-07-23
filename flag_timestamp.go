package cli

import (
	"time"
)

type TimestampFlag = FlagBase[time.Time, TimestampConfig, timestampValue]

type TimestampConfig struct {
	Timezone *time.Location

	Layouts []string
}

type timestampValue struct {
	timestamp  *time.Time
	hasBeenSet bool
	layouts    []string
	location   *time.Location
}

var _ ValueCreator[time.Time, TimestampConfig] = timestampValue{}

func (t timestampValue) Create(val time.Time, p *time.Time, c TimestampConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (t timestampValue) ToString(b time.Time) string { _ = "STUB: not implemented"; return "" }

func (t *timestampValue) Set(value string) error { _ = "STUB: not implemented"; return nil }

func (t *timestampValue) String() string { _ = "STUB: not implemented"; return "" }

func (t *timestampValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (cmd *Command) Timestamp(name string) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
