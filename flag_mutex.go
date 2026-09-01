package cli

type MutuallyExclusiveFlags struct {
	Flags [][]Flag

	Required bool

	Category string
}

func (grp MutuallyExclusiveFlags) check(_ *Command) error { _ = "STUB: not implemented"; return nil }

func (grp MutuallyExclusiveFlags) findSetFlag(startIdx int) (string, int, bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func (grp MutuallyExclusiveFlags) propagateCategory() { _ = "STUB: not implemented"; return }
