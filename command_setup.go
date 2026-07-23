package cli

func (cmd *Command) setupDefaults(osArgs []string) { _ = "STUB: not implemented"; return }

func (cmd *Command) setupCommandGraph() { _ = "STUB: not implemented"; return }

func (cmd *Command) setupSubcommand() { _ = "STUB: not implemented"; return }

func flagNamesInUse(flags []Flag, names []string) bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) hideHelp() bool { _ = "STUB: not implemented"; return false }

func (cmd *Command) ensureHelp() { _ = "STUB: not implemented"; return }

func dropClashingAliases(aliases []string, userFlags []Flag, selfName string) []string {
	_ = "STUB: not implemented"
	return nil
}
