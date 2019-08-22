package main

import (
	"fmt"
	"github.com/satori/go.uuid"

	"github.com/shevchenkobn/blog-backend/internal/di"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

func main() {
	// Start handling exits
	exitHandler := di.GetExitHandler()
	exitHandler.StartListeningToSignals()
	defer exitHandler.RecoverOrExit()

	posts, comments := di.GetPostRepository(), di.GetCommentRepository()
	fmt.Println(posts.GetAll())
	postId, err := uuid.FromString("8ae6f245-be64-4500-a6e1-85bc9ef43ee1")
	if err != nil {
		fmt.Println("uuid", err)
	}
	fmt.Println(posts.DeleteOne(postId, true))
	p, err := posts.CreateOne(&models.PostSeed{
		PostId: postId,
		AuthorName: "author",
		Content: "content",
	})
	fmt.Println(p, err)
	commentId, err := uuid.FromString("8ae6f245-be64-4500-a6e1-85bc9ef43e44")
	fmt.Println(comments.DeleteOne(commentId, true))
	fmt.Println(comments.CreateOne(&models.CommentSeed{
		CommentId: commentId,
		Post: p,
		Content: "comment",
		AuthorName: "authorCO",
	}))
	fmt.Println(posts.GetAll())
	fmt.Println(comments.GetAllForPost(postId))
}
