package typbuild

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func (b *Build) run(ctx context.Context, c *Context, args []string) (err error) {
	if err = b.buildProject(ctx, c); err != nil {
		return
	}
	log.Info("Run the application")
	cmd := exec.CommandContext(ctx, fmt.Sprintf("%s/%s", c.Bin, c.Name), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
