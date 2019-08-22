package pg

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
	commentm "github.com/shevchenkobn/blog-backend/internal/repository/model/comment"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/util"
	"time"
)

type comment struct {
	CommentIdField uuid.UUID `sql:"comment_id,pk,use_zero"`
	AuthorNameField string `sql:"author_name,notnull"`
	PostField *post
	ContentField string `sql:"content,notnull"`
	CommentedAtField time.Time `sql:"posted_at,default:(now() at time zone 'utc'),"` // FIXME: change to now
}
func (p *comment) CommentId() uuid.UUID {
	return p.CommentIdField
}
func (p *comment) AuthorName() string {
	return p.AuthorNameField
}
func (p *comment) SetAuthorName(author string) {
	p.AuthorNameField = author
}
func (p *comment) Post() models.Post {
	return p.PostField
}
func (p *comment) Content() string {
	return p.AuthorNameField
}
func (p *comment) SetContent(content string) {
	p.ContentField = content
}
func (p *comment) CommentedAt() time.Time {
	return p.CommentedAtField
}
func newComment(seed models.CommentSeed) (*comment, error) {
	c := new(comment)
	if seed.AuthorName == "" {
		return nil, &repository.ModelError{Code: commentm.AuthorNameRequired}
	}
	if seed.Post == nil {
		return nil, &repository.ModelError{Code: commentm.PostRequired}
	}
	post, ok := seed.Post.(*post)
	if !ok {
		return nil, &repository.ModelError{Code: commentm.PostInvalidType}
	}
	if seed.Content == "" {
		return nil, &repository.ModelError{Code: commentm.ContentRequired}
	}

	if c.CommentIdField == util.ZeroUuid {
		c.CommentIdField = uuid.NewV4()
	}
	c.AuthorNameField = seed.AuthorName
	c.PostField = post
	c.ContentField = seed.Content
	if c.CommentedAtField == util.ZeroTime {
		c.CommentedAtField = time.Now()
	}
	return c, nil
}

type CommentRepository struct {
	db *db.PostgreDB
}
func (r *CommentRepository) GetAllForPost(postId uuid.UUID) ([]models.Comment, error) {
	var post = &post{PostIdField: postId}
	err := r.db.Db().Model(post).Column("_").Relation("Comments").Select()
	if err != nil {
		return nil, err
	}
	//var commentsByInterface = make([]commentm.Comment, len(comments))
	//for i, c := range post.Comments() {
	//	commentsByInterface[i] = &c
	//}
	return post.Comments(), err
}
func NewCommentRepository(db *db.PostgreDB) *CommentRepository {
	r := new(CommentRepository)
	r.db = db
	// ensure table
	c := &comment{CommentIdField: util.ZeroUuid}
	err := r.db.Db().Select(&c)
	if err != nil {
		panic(err)
	}
	return r
}
