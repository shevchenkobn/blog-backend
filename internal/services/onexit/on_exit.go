package onexit

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

// readonly
var exitSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP}

func GetExitSignals() []os.Signal {
	return exitSignals
}

type ExitHandler struct {
	registeredCallbacks  []types.ExitHandlerCallback
	logger               *logger.Logger
	signals              []os.Signal
	isWindows            bool
	sigs                 chan os.Signal
	isListeningToSignals bool
	isRecovering         bool
}

func (handler *ExitHandler) IsListeningToSignals() bool {
	return handler.isListeningToSignals
}
func (handler *ExitHandler) IsRecovering() bool {
	return handler.isRecovering
}
func (handler *ExitHandler) SetRecovering(value bool) {
	handler.isRecovering = value
}

func NewExitHandler(l *logger.Logger) types.ExitHandler {
	handler := new(ExitHandler)
	handler.isWindows = runtime.GOOS == "windows"
	handler.logger = l
	handler.signals = GetExitSignals()
	handler.registeredCallbacks = make([]types.ExitHandlerCallback, 0, 4)
	handler.isListeningToSignals = false
	handler.isRecovering = true
	return handler
}

func (handler *ExitHandler) StartListeningToSignals() bool {
	if handler.isListeningToSignals {
		return false
	}
	handler.sigs = make(chan os.Signal)
	if !handler.isWindows {
		signal.Notify(handler.sigs, handler.signals...)
	} else {
		signal.Notify(handler.sigs, os.Interrupt)
	}
	go func() {
		sig := <-handler.sigs
		handler.logger.Printf("Got signal %v. Tidying up and exiting...", sig)
		handler.runExitCallbacks(true)
	}()
	handler.isListeningToSignals = true
	return true
}

func (handler *ExitHandler) StopListeningToSignals() bool {
	if !handler.isListeningToSignals {
		return false
	}
	close(handler.sigs)
	return true
}

func (handler *ExitHandler) RecoverOrExit() {
	if !handler.isRecovering {
		handler.runExitCallbacks(true)
		os.Exit(0)
		return
	}
	if err := recover(); err != nil {
		handler.logger.Errorf("Panic handled! %v. Tidying up and panicking again...", err)
		handler.runExitCallbacks(false)
		panic(err)
	}
}

func (handler *ExitHandler) AddCallback(callback types.ExitHandlerCallback) {
	handler.registeredCallbacks = append(handler.registeredCallbacks, callback)
}

func (handler *ExitHandler) RemoveCallback(callback types.ExitHandlerCallback) {
	handler.registeredCallbacks = util.RemoveCallbackOnOrder(handler.registeredCallbacks, callback)
}

func (handler *ExitHandler) runExitCallbacks(shouldExit bool) {
	for _, callback := range handler.registeredCallbacks {
		callback()
	}
	if shouldExit {
		os.Exit(0)
	}
}
