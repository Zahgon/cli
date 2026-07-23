package cli

type (
	StringMap     = MapBase[string, StringConfig, stringValue]
	StringMapFlag = FlagBase[map[string]string, StringConfig, StringMap]
)

var NewStringMap = NewMapBase[string, StringConfig, stringValue]

func (cmd *Command) StringMap(name string) map[string]string { _ = "STUB: not implemented"; return nil }
