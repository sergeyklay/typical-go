package typbuild

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func (b *Build) clean(ctx context.Context, c *Context) error {
	removeAllFile(c.Bin)
	removeAllFile(c.Temp)
	removeFile(".env") // TODO: configuration clean
	// removeFile(typenv.GeneratedConstructor) // TODO: app clean
	return nil
}

func removeFile(name string) {
	log.Infof("Remove: %s", name)
	if err := os.Remove(name); err != nil {
		log.Error(err.Error())
	}
}

func removeAllFile(path string) {
	log.Infof("Remove All: %s", path)
	if err := os.RemoveAll(path); err != nil {
		log.Error(err.Error())
	}
}
