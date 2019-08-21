package main

import (
	"fmt"
	"time"

	"github.com/shevchenkobn/blog-backend/internal/di"
)

func main() {
	// Start handling exits
	exitHandler := di.GetExitHandler()
	exitHandler.StartListeningToSignals()
	defer exitHandler.Recover()

	connection := di.GetPostgreDB()
	time.Sleep(5000)
	fmt.Println(connection)
	select {}
}
