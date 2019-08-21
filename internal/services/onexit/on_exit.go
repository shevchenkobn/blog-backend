package onexit

import (
	"os"
	"os/signal"
	"runtime"

	"../../util"
	"../logger"
)

type Callback func()
type ExitHandler struct {
	registeredCallbacks []Callback
	isWindows bool
	logger *logger.Logger
	sigs <-chan os.Signal
}

func NewExitHandler(l *logger.Logger, signals []os.Signal) *ExitHandler {
	handler := new(ExitHandler)
	handler.isWindows = runtime.GOOS == "windows"
	handler.logger = l

	sigs := make(chan os.Signal)
	handler.sigs = sigs
	if !handler.isWindows {
		signal.Notify(sigs, signals...)
		go func() {
			sig := <-handler.sigs
			handler.logger.Printf("Got signal %s", sig)
			handler.runExitCallbacks()
		}()
	} else {
		signal.Notify(sigs, os.Interrupt)
	}
	handler.registeredCallbacks = make([]Callback, 4, 4)
	return handler
}

func (handler *ExitHandler) AddCallback(callback Callback) {
	handler.registeredCallbacks = append(handler.registeredCallbacks, callback)
}

func (handler *ExitHandler) RemoveCallback(callback Callback) {
	handler.registeredCallbacks = util.RemoveCallbackOnOrder(handler.registeredCallbacks, callback)
}

func (handler *ExitHandler) runExitCallbacks() {
	for _, callback := range handler.registeredCallbacks {
		callback()
	}
}
