package cli

type CommandCategories interface {
	AddCommand(category string, command *Command)

	Categories() []CommandCategory
}

type commandCategories []*commandCategory

func newCommandCategories() CommandCategories {
	_ = "STUB: not implemented"
	return *new(CommandCategories)
}

func (c *commandCategories) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c *commandCategories) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *commandCategories) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (c *commandCategories) AddCommand(category string, command *Command) {
	_ = "STUB: not implemented"
	return
}

func (c *commandCategories) Categories() []CommandCategory { _ = "STUB: not implemented"; return nil }

type CommandCategory interface {
	Name() string

	VisibleCommands() []*Command
}

type commandCategory struct {
	name     string
	commands []*Command
}

func (c *commandCategory) Name() string { _ = "STUB: not implemented"; return "" }

func (c *commandCategory) VisibleCommands() []*Command { _ = "STUB: not implemented"; return nil }

type FlagCategories interface {
	AddFlag(category string, fl Flag)

	VisibleCategories() []VisibleFlagCategory
}

type defaultFlagCategories struct {
	m map[string]*defaultVisibleFlagCategory
}

func newFlagCategories() FlagCategories { _ = "STUB: not implemented"; return *new(FlagCategories) }

func newFlagCategoriesFromFlags(fs []Flag) FlagCategories {
	_ = "STUB: not implemented"
	return *new(FlagCategories)
}

func (f *defaultFlagCategories) AddFlag(category string, fl Flag) {
	_ = "STUB: not implemented"
	return
}

func (f *defaultFlagCategories) VisibleCategories() []VisibleFlagCategory {
	_ = "STUB: not implemented"
	return nil
}

type VisibleFlagCategory interface {
	Name() string

	Flags() []Flag
}

type defaultVisibleFlagCategory struct {
	name string
	m    map[string]Flag
}

func (fc *defaultVisibleFlagCategory) Name() string { _ = "STUB: not implemented"; return "" }

func (fc *defaultVisibleFlagCategory) Flags() []Flag { _ = "STUB: not implemented"; return nil }
