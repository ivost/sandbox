package api

func (s *Server) setupRoutes() {
	s.router.Handle("/version", s.handleVersion()).Methods("GET")
}
