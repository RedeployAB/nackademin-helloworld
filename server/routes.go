package server

// routes sets up the routes for the HTTP server.
func (s server) routes() {
	s.router.Handle("/", s.helloHandler())
}
