package action

import (
	"context"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/palantir/stacktrace"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v2"
)

func GeneratePlugin(c *cli.Context) error {
	if c.NArg() != 1 {
		return stacktrace.NewError("requires exactly 1 arg")
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return stacktrace.Propagate(err, "failed to create docker client from env")
	}

	absContextDir, err := validateContextDir(".")
	if err != nil {
		return stacktrace.Propagate(err, "failed to validate context directory")
	}

	createCtx, err := archive.TarWithOptions(absContextDir, &archive.TarOptions{
		Compression: archive.Uncompressed,
	})
	if err != nil {
		return stacktrace.Propagate(err, "failed to create docker client from env")
	}

	name := c.Args().Get(0)
	err = cli.PluginCreate(ctx, createCtx, types.PluginCreateOptions{
		RepoName: name,
	})
	if err != nil {
		return stacktrace.Propagate(err, "failed to create plugin '%s'", name)
	}

	return nil
}

func validateContextDir(contextDir string) (string, error) {
	absContextDir, err := filepath.Abs(contextDir)
	if err != nil {
		return "", err
	}
	stat, err := os.Lstat(absContextDir)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", errors.Errorf("context must be a directory")
	}

	return absContextDir, nil
}
