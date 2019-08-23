package pg

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/types"
	"time"

	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
	postm "github.com/shevchenkobn/blog-backend/internal/repository/model/post"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

type post struct {
	PostIdField     uuid.UUID `sql:"post_id,type:uuid,pk,use_zero" json:"postId"`
	AuthorNameField string    `sql:"author_name,notnull" json:"authorName"`
	ContentField    string    `sql:"content,notnull" json:"content"`
	PostedAtField   time.Time `sql:"posted_at,default:(now() at time zone 'utc')" json:"postedAt"`
	CommentsField   []comment `pg:"fk:parent_post_id" json:"-"`
}
const post_PK = "post_id"
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
func (p *post) Comments() []models.Comment {
	return toInterface(p.CommentsField)
}
func newPost(seed *models.PostSeed) (*post, error) {
	p := new(post)
	if seed.AuthorName == "" {
		return nil, &repository.ModelError{Code: postm.AuthorNameRequired}
	}
	if seed.Content == "" {
		return nil, &repository.ModelError{Code: postm.ContentRequired}
	}

	if seed.PostId == util.ZeroUuid {
		p.PostIdField = uuid.NewV4()
	} else {
		p.PostIdField = seed.PostId
	}
	p.AuthorNameField = seed.AuthorName
	p.ContentField = seed.Content
	if seed.PostedAt == util.ZeroTime {
		p.PostedAtField = time.Now()
	}
	p.CommentsField = make([]comment, 0, 1)
	return p, nil
}
func PostToJson(post models.Post) ([]byte, error) {
	return json.Marshal(post)
}
func PostsToJson(posts []models.Post) ([]byte, error) {
	return json.Marshal(posts)
}
var zeroPost = &post{}

type PostRepository struct {
	db *db.PostgreDB
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

func (r *PostRepository) CreateOne(post *models.PostSeed) (models.Post, error) {
	p, err := newPost(post)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Db().Model(p).Returning("*").Insert()
	if err != nil {
		if err, ok := err.(pg.Error); ok {
			if err.Field('C') == "23505" {
				return nil, types.NewLogicError(types.ErrorPostDuplicateId)
			}
		}
		return nil, err
	}
	return p, nil
}

func (r *PostRepository) DeleteOne(postId uuid.UUID, returning bool) (models.Post, error) {
	c := &post{}
	q := r.db.Db().Model(c).Where("post_id = ?", postId)
	if returning {
		q = q.Returning("*")
	}
	_, err := q.Delete()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, types.NewLogicError(types.ErrorNotFound)
		}
		return nil, err
	}
	if returning {
		return c, nil
	} else {
		return nil, nil
	}
}
