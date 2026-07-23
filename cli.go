package cli

import (
	"os"
)

var isTracingOn = os.Getenv("URFAVE_CLI_TRACING") == "on"

func tracef(format string, a ...any) { _ = "STUB: not implemented"; return }
