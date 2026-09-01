package cli

import (
	"time"
)

type DurationFlag = FlagBase[time.Duration, NoConfig, durationValue]

type durationValue time.Duration

func (d durationValue) Create(val time.Duration, p *time.Duration, c NoConfig) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (d durationValue) ToString(val time.Duration) string { _ = "STUB: not implemented"; return "" }

func (d *durationValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (d *durationValue) Get() any { _ = "STUB: not implemented"; return *new(any) }

func (d *durationValue) String() string { _ = "STUB: not implemented"; return "" }

func (cmd *Command) Duration(name string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
