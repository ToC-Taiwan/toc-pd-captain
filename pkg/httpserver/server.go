// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout       = 5 * time.Second
	_defaultReadHeaderTimeout = 5 * time.Second
	_defaultWriteTimeout      = 5 * time.Minute
	_defaultAddr              = ":80"
	_defaultShutdownTimeout   = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, opts ...Option) *Server {
	s := &Server{
		server: &http.Server{
			Handler:           handler,
			ReadHeaderTimeout: _defaultReadHeaderTimeout,
			ReadTimeout:       _defaultReadTimeout,
			WriteTimeout:      _defaultWriteTimeout,
			Addr:              _defaultAddr,
		},
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()
	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
