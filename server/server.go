package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// logger is the interface that wraps around method Printf.
type logger interface {
	Printf(format string, v ...any)
}

// server contains the http server and router.
type server struct {
	httpServer *http.Server
	router     *http.ServeMux
	log        logger
}

// Options contains the server options.
type Options struct {
	Router *http.ServeMux
	Log    logger
	Port   string
	Host   string
}

// New returns a new server.
func New(options ...Options) *server {
	opts := Options{}
	for _, option := range options {
		opts = option
	}

	if opts.Router == nil {
		opts.Router = http.NewServeMux()
	}

	if opts.Log == nil {
		opts.Log = log.Default()
	}

	if len(opts.Port) == 0 {
		opts.Port = "8080"
	}

	srv := &http.Server{
		Addr:         opts.Host + ":" + opts.Port,
		Handler:      opts.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &server{
		httpServer: srv,
		router:     opts.Router,
		log:        opts.Log,
	}
}

// Start the server.
func (s server) Start() {
	s.routes()
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Printf("Could not start server: %v\n", err)
		}
	}()
	s.log.Printf("Server started, listening on %s.\n", s.httpServer.Addr)

	sig, err := s.stop()
	if err != nil {
		s.log.Printf("Could not stop server: %v\n", err)
	}
	s.log.Printf("Server stopped, reason: %s.\n", sig.String())
}

// stop the server gracefully.
func (s server) stop() (os.Signal, error) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.httpServer.SetKeepAlivesEnabled(false)
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return nil, err
	}
	return sig, nil
}
