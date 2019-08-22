package main

import (
	"fmt"
	"github.com/shevchenkobn/blog-backend/internal/di"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

func main() {
	// Start handling exits
	exitHandler := di.GetExitHandler()
	exitHandler.StartListeningToSignals()
	//defer exitHandler.RecoverOrExit()

	posts, comments := di.GetPostRepository(), di.GetCommentRepository()
	fmt.Println(posts.GetAll())
	fmt.Println(comments.GetAllForPost(util.ZeroUuid))
}
