package types

import (
	"encoding/json"
	"net/http"

	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

type ErrorCode string
const (
	ErrorPostDuplicateId = "POST_ID_DUPLICATE"

	ErrorCommentDuplicateId   = "COMMENT_ID_DUPLICATE"
	ErrorCommentInvalidBlogId = "COMMENT_BLOG_ID_INVALID"

	ErrorNotFound = "NOT_FOUND"

	ErrorServer = "SERVER"
)

type LogicError interface {
	Code() ErrorCode
	Message() string
	ToJson() ([]byte, error)
}

func NewLogicError(code ErrorCode) LogicError {
	return &logicError{CodeField: code}
}

func NewLogicErrorWithMessage(code ErrorCode, message string) LogicError {
	return &logicError{CodeField: code, MessageField: message}
}

type logicError struct {
	CodeField ErrorCode `json:"code"`
	MessageField string `json:"message,omitempty"`
}

func (e *logicError) Code() ErrorCode {
	return e.CodeField
}

func (e *logicError) Message() string {
	return e.MessageField
}

func (e *logicError) ToJson() ([]byte, error) {
	return json.Marshal(e)
}

func SendLogicError(w http.ResponseWriter, logger *logger.Logger, code int, err LogicError) {
	bytes, e := err.ToJson()
	if e != nil {
		util.SendResponse(w, logger, http.StatusInternalServerError, GetServerFailureError())
	} else {
		util.SendResponse(w, logger, code, bytes)
	}
}

func GetServerFailureError() []byte {
	return []byte("{ code: \"" + ErrorServer + "\" }")
}
