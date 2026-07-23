package cli

type (
	StringSlice     = SliceBase[string, StringConfig, stringValue]
	StringSliceFlag = FlagBase[[]string, StringConfig, StringSlice]
)

var NewStringSlice = NewSliceBase[string, StringConfig, stringValue]

func (cmd *Command) StringSlice(name string) []string { _ = "STUB: not implemented"; return nil }
