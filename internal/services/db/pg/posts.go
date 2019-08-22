package pg

import (
	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
	postm "github.com/shevchenkobn/blog-backend/internal/repository/model/post"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/util"
	"time"
)

type post struct {
	PostIdField     uuid.UUID `sql:"post_id,type:uuid,pk,use_zero"`
	AuthorNameField string    `sql:"author_name,notnull"`
	ContentField    string    `sql:"content,notnull"`
	PostedAtField   time.Time `sql:"posted_at,default:(now() at time zone 'utc')"` // FIXME: change to now
	Comments        []comment `sql:"on_delete:CASCADE"`
}
func (p *post) PostId() uuid.UUID {
	return p.PostIdField
}
func (p *post) AuthorName() string {
	return p.AuthorNameField
}
func (p *post) SetAuthorName(author string) {
	p.AuthorNameField = author
}
func (p *post) Content() string {
	return p.AuthorNameField
}
func (p *post) SetContent(content string) {
	p.ContentField = content
}
func (p *post) PostedAt() time.Time {
	return p.PostedAtField
}
func (p *post) GetComments() []models.Comment {
	return toInterface(p.Comments)
}
func newPost(seed models.PostSeed) (*post, error) {
	p := new(post)
	if seed.AuthorName == "" {
		return nil, &repository.ModelError{Code: postm.AuthorNameRequired}
	}
	if seed.Content == "" {
		return nil, &repository.ModelError{Code: postm.ContentRequired}
	}

	if seed.PostId == util.ZeroUuid {
		p.PostIdField = uuid.NewV4()
	}
	p.AuthorNameField = seed.AuthorName
	p.ContentField = seed.Content
	if seed.PostedAt == util.ZeroTime {
		p.PostedAtField = time.Now()
	}
	p.Comments = make([]comment, 0, 1)
	return p, nil
}
var zeroPost = &post{}

type PostRepository struct {
	db *db.PostgreDB
}
func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []post
	err := r.db.Db().Model(&posts).Select()
	if err != nil {
		return nil, err
	}
	var postsByInterface = make([]models.Post, len(posts))
	for i, p := range posts {
		postsByInterface[i] = &p
	}
	return postsByInterface, err
}
func NewPostRepository(db *db.PostgreDB) *PostRepository {
	r := new(PostRepository)
	r.db = db
	err := r.db.Db().CreateTable(zeroPost, &orm.CreateTableOptions{
		FKConstraints: true,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
	return r
}
