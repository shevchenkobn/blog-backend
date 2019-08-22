package main

import (
	"github.com/shevchenkobn/blog-backend/internal/di"
)

func main() {
	// Start handling exits
	exitHandler := di.GetExitHandler()
	exitHandler.StartListeningToSignals()
	defer exitHandler.RecoverOrExit()

	di.GetServer().ListenAndWait()
}
