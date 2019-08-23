package openapi

import (
	"context"
	"fmt"
	"github.com/MNFGroup/openapimux"
	"github.com/shevchenkobn/blog-backend/handlers"
	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
	"net/http"
	"strconv"
	"sync"
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
		newErr := srv.Shutdown(context.TODO())
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

func NewServer(config config.Config, exitHandler types.ExitHandler, logger *logger.Logger) *Server {
	r, err := openapimux.NewRouter(config.OpenApi().ConfigPath())
	if err != nil {
		panic(err)
	}
	r.UseHandlers(map[string]http.Handler{
		"GetPosts": handlers.GetPosts{},
	})
	r.ErrorHandler = func(w http.ResponseWriter, r *http.Request, data string, code int) {
		w.WriteHeader(code)
		if code == http.StatusInternalServerError {
			fmt.Println("Fatal:", data)
			w.Write([]byte("Oops"))
		} else {
			w.Write([]byte(data))
		}
	}

	server := new(Server)
	server.config = config.Servers()
	server.r = r
	server.exitHandler = exitHandler
	server.logger = logger
	return server
}
