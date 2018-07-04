package chat

func (s *Server) initRoute() {
	s.mux.GET("/room/:room_id", s.connectRoom)
}
