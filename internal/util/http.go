package util

import (
	"net/http"

	"github.com/shevchenkobn/blog-backend/internal/services/logger"
)

func WriteSafely(w http.ResponseWriter, logger *logger.Logger, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		logger.Errorf("Error responding to request: %v", err)
	}
}

func SendResponse(w http.ResponseWriter, logger *logger.Logger, code int, message []byte) {
	w.WriteHeader(code)
	WriteSafely(w, logger, message)
}

