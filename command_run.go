package cli

import (
	"context"
)

type helpShownKey struct{}

func (cmd *Command) parseArgsFromStdin() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *Command) Run(ctx context.Context, osArgs []string) (deferErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *Command) run(ctx context.Context, osArgs []string) (_ context.Context, deferErr error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (cmd *Command) handleRequiredError(ctx context.Context, err error) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func commandChain(cmd *Command) []*Command { _ = "STUB: not implemented"; return nil }

func findArgValidator(cmd *Command) ArgValidatorFunc {
	_ = "STUB: not implemented"
	return *new(ArgValidatorFunc)
}

func runBefore(ctx context.Context, cmdChain []*Command) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
