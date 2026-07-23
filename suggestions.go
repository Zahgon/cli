package cli

const suggestDidYouMeanTemplate = "Did you mean %q?"

var (
	SuggestFlag               SuggestFlagFunc    = suggestFlag
	SuggestCommand            SuggestCommandFunc = suggestCommand
	SuggestDidYouMeanTemplate string             = suggestDidYouMeanTemplate
)

type SuggestFlagFunc func(flags []Flag, provided string, hideHelp bool) string

type SuggestCommandFunc func(commands []*Command, provided string) string

func jaroDistance(a, b string) float64 { _ = "STUB: not implemented"; return 0 }

func jaroWinkler(a, b string) float64 { _ = "STUB: not implemented"; return 0 }

func suggestFlag(flags []Flag, provided string, hideHelp bool) string {
	_ = "STUB: not implemented"
	return ""
}

func suggestCommand(commands []*Command, provided string) (suggestion string) {
	_ = "STUB: not implemented"
	return ""
}
