package openapi

import (
	"context"
	"github.com/MNFGroup/openapimux"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/middleware"

	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
)

type Server struct {
	config []config.ServerConfig
	r *openapimux.OpenAPIMux
	servers []*http.Server
	wg *sync.WaitGroup
	exitHandler types.ExitHandler
	exitCallback types.ExitHandlerCallback
	logger *logger.Logger
}
func (s *Server) ListenAndWait() {
	s.servers = make([]*http.Server, 0, len(s.config))
	s.wg = &sync.WaitGroup{}
	s.wg.Add(1)
	for _, c := range s.config {
		srv := &http.Server{Addr: c.Host() + ":" + strconv.Itoa(c.Port()), Handler: s.r}
		s.servers = append(s.servers, srv)
		go func() {
			s.logger.Printf("Listening to %s", srv.Addr)
			err := srv.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}()
	}
	s.exitCallback = func() {
		s.Close()
	}
	s.exitHandler.AddCallback(s.exitCallback)
	s.wg.Wait()
}
func (s *Server) Close() {
	var err error
	for _, srv := range s.servers {
		newErr := srv.Shutdown(context.Background())
		if newErr != nil {
			if err != nil {
				s.logger.Errorf("Error while stopping server: %v", err)
			}
		}
		newErr = err
	}
	s.exitHandler.RemoveCallback(s.exitCallback)
	if err != nil {
		panic(err)
	}
	s.wg.Done()
}

func NewServer(config config.Config, handlers map[string]http.Handler, exitHandler types.ExitHandler, logger *logger.Logger) *Server {
	r, err := openapimux.NewRouter(config.OpenApi().ConfigPath())
	if err != nil {
		panic(err)
	}
	r.DetailedError = true
	r.UseHandlers(handlers)
	r.UseMiddleware(
		ErrorHandler(logger),
		middleware.SetHeader(http.CanonicalHeaderKey("content-type"), "application/json"),
	)
	r.ErrorHandler = func(w http.ResponseWriter, r *http.Request, data string, code int) {
		switch data {
		case "Path not found":
			types.SendLogicError(w, logger, http.StatusNotFound, types.NewLogicError(types.ErrorNotFound))
		case "Handler not found":
			logger.Errorf("Handler not found: %s", r.URL.String())
			types.SendLogicError(w, logger, http.StatusInternalServerError, types.NewLogicError(types.ErrorServer))
		default:
			types.SendLogicError(w, logger, http.StatusBadRequest, types.NewLogicErrorWithMessage(types.ErrorServer, data))
		}
	}

	server := new(Server)
	server.config = config.Servers()
	server.r = r
	server.exitHandler = exitHandler
	server.logger = logger
	return server
}
