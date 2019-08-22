package pg

import (
	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
	commentm "github.com/shevchenkobn/blog-backend/internal/repository/model/comment"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

type comment struct {
	CommentIdField   uuid.UUID `sql:"comment_id,pk,type:uuid,use_zero"`
	AuthorNameField  string    `sql:"author_name,notnull"`
	PostId           uuid.UUID `sql:"parent_post_id,type:uuid,notnull,on_delete:CASCADE"`
	PostField        *post     `pg:"fk:parent_post_id"`
	ContentField     string    `sql:"content,notnull"`
	CommentedAtField time.Time `sql:"posted_at,default:(now() at time zone 'utc'),"` // FIXME: change to now
}
const commentPK = "comment_id"
const postFK = "parent_post_id"
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
func newComment(seed *models.CommentSeed) (*comment, error) {
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
	} else {
		c.CommentIdField = seed.CommentId
	}
	c.AuthorNameField = seed.AuthorName
	c.PostId = post.PostIdField
	c.ContentField = seed.Content
	if c.CommentedAtField == util.ZeroTime {
		c.CommentedAtField = time.Now()
	}
	return c, nil
}
var zeroComment = &comment{}

type CommentRepository struct {
	db *db.PostgreDB
}

func NewCommentRepository(db *db.PostgreDB) *CommentRepository {
	r := new(CommentRepository)
	r.db = db
	err := r.db.Db().CreateTable(zeroComment, &orm.CreateTableOptions{
		FKConstraints: true,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
	return r
}

func (r *CommentRepository) GetAllForPost(postId uuid.UUID) ([]models.Comment, error) {
	var comments []comment
	err := r.db.Db().Model(&comments).Where(postFK + " = ?", postId).Select()
	if err != nil {
		return nil, err
	}
	return toInterface(comments), nil
}

func (r *CommentRepository) CreateOne(comment *models.CommentSeed) (models.Comment, error) {
	c, err := newComment(comment)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Db().Model(c).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CommentRepository) DeleteOne(commentId uuid.UUID, returning bool) (models.Comment, error) {
	c := &comment{}
	q := r.db.Db().Model(c).Where("comment_id = ?", commentId)
	if returning {
		q = q.Returning("*")
	}
	_, err := q.Delete()
	if err != nil {
		return nil, err
	}
	if returning {
		return c, nil
	} else {
		return nil, nil
	}
}

func toInterface(comments []comment) []models.Comment {
	var commentsByInterface = make([]models.Comment, len(comments))
	for i, c := range comments {
		commentsByInterface[i] = &c
	}
	return commentsByInterface
}
