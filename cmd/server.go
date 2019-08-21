package main

import (
	"time"

	"../internal/di"
)

func main() {
	connection := di.GetPostgreDB()
	time.Sleep(1)
	connection.Close()
}
