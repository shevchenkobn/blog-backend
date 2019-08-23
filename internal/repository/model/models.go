package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Comment interface {
	CommentId() uuid.UUID
	AuthorName() string
	SetAuthorName(author string)
	PostId() uuid.UUID
	Post() Post
	Content() string
	SetContent(content string)
	CommentedAt() time.Time
}
type CommentSeed struct {
	CommentId   uuid.UUID
	AuthorName  string
	Post        Post
	Content     string
	CommentedAt time.Time
}

type Post interface {
	PostId() uuid.UUID
	AuthorName() string
	SetAuthorName(author string)
	Content() string
	SetContent(content string)
	PostedAt() time.Time
	Comments() []Comment // readonly, no copying for memory efficiency
}
type PostSeed struct {
	PostId     uuid.UUID
	AuthorName string
	Content    string
	PostedAt   time.Time
}
