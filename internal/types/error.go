package types

import (
	"encoding/json"
)

type ErrorCode string

const (
	ErrorOpenApi = "OPENAPI"

	ErrorPostDuplicateId = "POST_ID_DUPLICATE"
	ErrorPostNotFound    = "POST_NOT_FOUND"

	ErrorCommentDuplicateId   = "COMMENT_ID_DUPLICATE"
	ErrorCommentInvalidBlogId = "COMMENT_BLOG_ID_INVALID"
	ErrorCommentNotFound      = "COMMENT_NOT_FOUND"

	ErrorNotFound = "NOT_FOUND"

	ErrorServer = "SERVER"
)

type LogicError interface {
	error
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
	CodeField    ErrorCode `json:"code"`
	MessageField string    `json:"message,omitempty"`
}

func (e *logicError) Error() string {
	return "LogicError: " + string(e.Code())
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

func GetServerFailureError() []byte {
	return []byte("{ code: \"" + ErrorServer + "\" }")
}
