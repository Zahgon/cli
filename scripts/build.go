package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

const (
	badNewsEmoji      = "🚨"
	goodNewsEmoji     = "✨"
	checksPassedEmoji = "✅"

	gfmrunVersion = "v1.3.0"

	v3diffWarning = `
# The unified diff above indicates that the public API surface area
# has changed. If you feel that the changes are acceptable for the
# v3.x series, please run the following command to promote the
# current go docs:
#
#     make v3approve
#
`
)

func main() {
	topDir, err := func() (string, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if v, err := sh(ctx, "git", "rev-parse", "--show-toplevel"); err == nil {
			return strings.TrimSpace(v), nil
		}

		return os.Getwd()
	}()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.Command{
		Name:  "builder",
		Usage: "Do a thing for urfave/cli! (maybe build?)",
		Commands: []*cli.Command{
			{
				Name:   "vet",
				Action: topRunAction("go", "vet", "./..."),
			},
			{
				Name:   "test",
				Action: TestActionFunc,
			},
			{
				Name: "gfmrun",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "walk",
						Value: false,
						Usage: "Walk the specified directory and perform validation on all markdown files",
					},
				},
				Action: GfmrunActionFunc,
			},
			{
				Name:   "check-binary-size",
				Action: checkBinarySizeActionFunc,
			},
			{
				Name:   "generate",
				Action: GenerateActionFunc,
			},
			{
				Name:   "diffcheck",
				Action: DiffCheckActionFunc,
			},
			{
				Name:   "ensure-goimports",
				Action: EnsureGoimportsActionFunc,
			},
			{
				Name:   "ensure-gfmrun",
				Action: EnsureGfmrunActionFunc,
			},
			{
				Name:   "ensure-mkdocs",
				Action: EnsureMkdocsActionFunc,
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "upgrade-pip"},
				},
			},
			{
				Name:   "set-mkdocs-remote",
				Action: SetMkdocsRemoteActionFunc,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "github-token",
						Sources:  cli.EnvVars("MKDOCS_REMOTE_GITHUB_TOKEN"),
						Required: true,
					},
				},
			},
			{
				Name:   "deploy-mkdocs",
				Action: topRunAction("mkdocs", "gh-deploy", "--force"),
			},
			{
				Name:   "lint",
				Action: LintActionFunc,
			},
			{
				Name: "v3diff",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "color", Value: false},
				},
				Action: V3Diff,
			},
			{
				Name: "v3approve",
				Action: topRunAction(
					"cp",
					"-v",
					"godoc-current.txt",
					filepath.Join("testdata", "godoc-v3.x.txt"),
				),
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tags",
				Usage: "set build tags",
			},
			&cli.StringFlag{
				Name:  "top-dir",
				Value: topDir,
			},
			&cli.StringSliceFlag{
				Name:  "packages",
				Value: []string{"cli", "scripts"},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func sh(ctx context.Context, exe string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func topRunAction(arg string, args ...string) cli.ActionFunc {
	_ = "STUB: not implemented"
	return *new(cli.ActionFunc)
}

func runCmd(ctx context.Context, arg string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadFile(src, dest string, dirPerm, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func VetActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func TestActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func testCleanup(packages []string) error { _ = "STUB: not implemented"; return nil }

func GfmrunActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func checkBinarySizeActionFunc(ctx context.Context, cmd *cli.Command) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func GenerateActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func DiffCheckActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureGoimportsActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureGfmrunActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureMkdocsActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func SetMkdocsRemoteActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func LintActionFunc(ctx context.Context, cmd *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func V3Diff(ctx context.Context, cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }

func getSize(ctx context.Context, sourcePath, builtPath, tags string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
