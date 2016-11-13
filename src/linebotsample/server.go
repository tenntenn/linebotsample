package linebot

import "net/http"

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}

	s.initHandler()

	return s
}

func (s *Server) initHandler() {
	s.mux.HandleFunc("/hook/", hookHandler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
