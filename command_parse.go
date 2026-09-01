package cli

const (
	providedButNotDefinedErrMsg = "flag provided but not defined: -"
	argumentNotProvidedErrMsg   = "flag needs an argument: "
)

func flagFromError(err error) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *Command) parseFlags(args Args) (Args, error) {
	_ = "STUB: not implemented"
	return *new(Args), nil
}
