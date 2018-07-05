package chat

import (
	"log"
)

type Server struct {
	users      map[UserID]*User
	publisher  Publisher
	subscriber Subscriber
	userEnter  chan *UserEnterRequest
	userLeave  chan *UserLeaveRequest
}

func NewServer(publisher Publisher, subscriber Subscriber) *Server {
	server := &Server{
		users:      make(map[UserID]*User),
		publisher:  publisher,
		subscriber: subscriber,
		userEnter:  make(chan *UserEnterRequest),
		userLeave:  make(chan *UserLeaveRequest),
	}
	return server
}

func (s *Server) OnUserEnter(req *UserEnterRequest) {
	log.Println("User entering room", req.RoomID, "with user id", req.UserID)
	s.userEnter <- req
}

func (s *Server) Run() {
	for {
		select {
		case req := <-s.userEnter:
			log.Println("User enter")
			s.handleUserEnter(req)
		case req := <-s.userLeave:
			log.Println("User leaving")
			s.handleUserLeave(req)
		}
	}
}

func (s *Server) handleUserEnter(req *UserEnterRequest) {
	usr := NewUser(s, UserID("antony"), RoomID(req.RoomID), "Antony", s.publisher, s.subscriber, req.Conn)
	s.users[usr.UserID] = usr
	usr.Run()

	log.Println("New user connected", usr)
	log.Println(len(s.users))
}

func (s *Server) handleUserLeave(req *UserLeaveRequest) {
	log.Println("Removing user")
	log.Println(req.User)
	log.Println(len(s.users))
	delete(s.users, req.User.UserID)
	log.Println(len(s.users))
}
