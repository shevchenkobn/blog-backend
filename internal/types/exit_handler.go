package types

type ExitHandlerCallback func()

type ExitHandler interface {
	IsListeningToSignals() bool
	IsRecovering() bool
	SetRecovering(value bool)

	StartListeningToSignals() bool
	StopListeningToSignals() bool
	Recover()
	AddCallback(callback ExitHandlerCallback)
	RemoveCallback(callback ExitHandlerCallback)
}
