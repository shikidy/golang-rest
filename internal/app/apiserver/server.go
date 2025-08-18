package apiserver

import (
	"net/http"

	"github.com/shikidy/golang-rest/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *http.ServeMux
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: http.NewServeMux(),
		store:  store,
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("POST /users", s.handleUsersCreate())
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
