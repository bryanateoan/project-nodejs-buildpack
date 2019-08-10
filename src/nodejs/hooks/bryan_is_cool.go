package hooks

import (
	"errors"
	"github.com/cloudfoundry/libbuildpack"
	"os"
)

type BryanHook struct {
	libbuildpack.DefaultHook
	Message string
	Log *libbuildpack.Logger
}

func init() {
	logger := libbuildpack.NewLogger(os.Stdout)
	libbuildpack.AddHook(BryanHook{
		Message: "Bryan is cool!",
		Log: logger,
	})
}

func (h BryanHook) AfterCompile(stager *libbuildpack.Stager) error {
	err := h.sayMessage()
	return err
}

func (h BryanHook) sayMessage() error {
	if h.Message == "" {
		h.Log.Error("Message is empty. Failing build...")
		return errors.New("no message to print")
	} else {
		h.Log.Protip(h.Message,"")
	}
	return nil
	
}
