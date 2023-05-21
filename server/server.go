package server

import (
	"io"
	"log"
	"net/http"

	"github.com/sivaramsajeev/rlimiter/configs"
	"github.com/sivaramsajeev/rlimiter/ratelimiter"
)

type Server struct {
	RateLimiter *ratelimiter.RateLimiter
}

func (s *Server) Serve() {
	s.RateLimiter.StartRateLimiter()
	s.serveHttp()
}

func (s *Server) serveHttp() {
	http.HandleFunc("/", s.getRoot())
	log.Println("Server starting on port:", configs.HTTP_PORT)
	err := http.ListenAndServe(configs.HTTP_PORT, nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) getRoot() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.RateLimiter.Bucket.ReduceByOne(); err != nil {
			io.WriteString(w, configs.RATE_LIMITED_MSG)
			return
		}
		io.WriteString(w, configs.NORMAL_MSG)
	}
}
