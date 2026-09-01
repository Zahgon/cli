package cli

import (
	"fmt"
	"io"
	"os"
)

var OsExiter = os.Exit

var ErrWriter io.Writer = os.Stderr

type MultiError interface {
	error
	Errors() []error
}

func newMultiError(err ...error) MultiError { _ = "STUB: not implemented"; return *new(MultiError) }

type multiError []error

func (m *multiError) Error() string { _ = "STUB: not implemented"; return "" }

func (m *multiError) Errors() []error { _ = "STUB: not implemented"; return nil }

type requiredFlagsErr interface {
	error
}

type errRequiredFlags struct {
	missingFlags []string
}

func (e *errRequiredFlags) Error() string { _ = "STUB: not implemented"; return "" }

type requiredArgumentsErr interface {
	error
}

type errRequiredArguments struct {
	missingArguments []string
}

func (e *errRequiredArguments) Error() string { _ = "STUB: not implemented"; return "" }

type mutuallyExclusiveGroup struct {
	flag1Name string
	flag2Name string
}

func (e *mutuallyExclusiveGroup) Error() string { _ = "STUB: not implemented"; return "" }

type mutuallyExclusiveGroupRequiredFlag struct {
	flags *MutuallyExclusiveFlags
}

func (e *mutuallyExclusiveGroupRequiredFlag) Error() string { _ = "STUB: not implemented"; return "" }

type ErrorFormatter interface {
	Format(s fmt.State, verb rune)
}

type ExitCoder interface {
	error
	ExitCode() int
}

type exitError struct {
	exitCode int
	err      error
}

func Exit(message any, exitCode int) ExitCoder { _ = "STUB: not implemented"; return *new(ExitCoder) }

func (ee *exitError) Error() string { _ = "STUB: not implemented"; return "" }

func (ee *exitError) ExitCode() int { _ = "STUB: not implemented"; return 0 }

func HandleExitCoder(err error) { _ = "STUB: not implemented"; return }

func handleMultiError(multiErr MultiError) int { _ = "STUB: not implemented"; return 0 }
