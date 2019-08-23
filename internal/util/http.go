package util

import (
	"github.com/MNFGroup/openapimux"
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/types"
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

func SendLogicError(w http.ResponseWriter, logger *logger.Logger, code int, err types.LogicError) {
	bytes, e := err.ToJson()
	if e != nil {
		SendResponse(w, logger, http.StatusInternalServerError, types.GetServerFailureError())
	} else {
		SendResponse(w, logger, code, bytes)
	}
}

func GetUuidPathParam(r *http.Request, name string) uuid.UUID {
	str := openapimux.PathParam(r, name)
	id, err := uuid.FromString(str)
	if err != nil {
		panic(err)
	}
	return id
}
