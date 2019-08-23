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
	CommentId   uuid.UUID `json:"commentId,omitempty"`
	AuthorName  string    `json:"authorName"`
	PostId      uuid.UUID `json:"-"`
	Post        Post      `json:"-"`
	Content     string    `json:"content"`
	CommentedAt time.Time `json:"commentedAt,omitempty"`
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
	PostId     uuid.UUID `json:"postId,omitempty"`
	AuthorName string    `json:"authorName"`
	Content    string    `json:"content"`
	PostedAt   time.Time `json:"postedAt,omitempty"`
}
