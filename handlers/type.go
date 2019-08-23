package handlers

import (
	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model/post"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"net/http"
	"reflect"
)

type HttpHandlers struct {
	postRepository repository.Posts
	postsToJson post.SliceJsonEncoder
	logger *logger.Logger
}

const mapperFunctionName = "ToMap"
func (ctx *HttpHandlers) ToMap() map[string]http.Handler {
	handlersType := reflect.ValueOf(ctx)
	methodCount := handlersType.NumMethod() - 1 // exclude current method
	ctx.logger.Printf("Handler count: %d", methodCount)
	m := make(map[string]http.Handler, methodCount)
	for i := 0; i < handlersType.NumMethod(); i += 1 {
		method := handlersType.Method(i)
		name := method.Name
		if name == mapperFunctionName {
			continue
		}
		m[name] =
	}

}

func New(postRepository repository.Posts, postsToJson post.SliceJsonEncoder, logger *logger.Logger) *HttpHandlers {
	h := new(HttpHandlers)
	h.postRepository = postRepository
	h.postsToJson = postsToJson
	h.logger = logger
	return h
}
