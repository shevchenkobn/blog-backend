package openapi

import (
	"net/http"

	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

type errorHandler struct {
	handler http.Handler
	logger *logger.Logger
}

func (handler *errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer handler.recover(w, r)
	handler.handler.ServeHTTP(w, r)
}

func (handler *errorHandler) recover(w http.ResponseWriter, r *http.Request) {
	err := recover()

	if err, ok := err.(types.LogicError); ok {
		var code int
		switch err.Code() {
		case types.ErrorServer:
			code = http.StatusBadRequest
			handler.logger.Errorf("Server error %v", err)
		default:
			code = http.StatusInternalServerError
		}
		types.SendLogicError(w, handler.logger, code, err)
		return
	}

	handler.logger.Errorf("Unexpected error %v", err)
	util.SendResponse(w, handler.logger, http.StatusInternalServerError, []byte(err.(string)))
}

func ErrorHandler(logger *logger.Logger) func(http.Handler) http.Handler {
	return func (handler http.Handler) http.Handler {
		return &errorHandler{handler: handler, logger: logger}
	}
}

