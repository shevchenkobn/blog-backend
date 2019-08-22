package main

import (
	"fmt"
	"github.com/satori/go.uuid"

	"github.com/shevchenkobn/blog-backend/internal/di"
)

func main() {
	// Start handling exits
	exitHandler := di.GetExitHandler()
	exitHandler.StartListeningToSignals()
	defer exitHandler.RecoverOrExit()

	posts, comments := di.GetPostRepository(), di.GetCommentRepository()
	fmt.Println(posts.GetAll())
	uuid, err := uuid.FromString("8ae6f245-be64-4500-a6e1-85bc9ef43ee1")
	if err != nil {
		fmt.Println("uuid", err)
	}
	fmt.Println(comments.GetAllForPost(uuid))
}
