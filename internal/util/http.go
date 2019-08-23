package util

import (
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"net/http"
)

func WriteSafely(w http.ResponseWriter, logger *logger.Logger, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		logger.Errorf("Error responding to request: %v", err)
	}
}
