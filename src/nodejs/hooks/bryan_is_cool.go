package hooks

import (
	
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
	logger.Protip("hello", "hi")
	libbuildpack.AddHook(BryanHook{
		Message: "Bryan is cool!",
		Log: logger,
		})
}

func (h BryanHook) AfterCompile(stager *libbuildpack.Stager) error {
	h.SayMessage()
	return nil
}

func (h BryanHook) SayMessage() {
	h.Log.Protip(h.Message,"")
}