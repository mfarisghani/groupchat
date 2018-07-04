package chat

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

type Server struct {
	users      map[UserID]*User
	port       string
	mux        *httprouter.Router
	upgrader   websocket.Upgrader
	publisher  Publisher
	subscriber Subscriber
	removeUser chan *User
}

func NewServer(port string, publisher Publisher, subscriber Subscriber) *Server {
	server := &Server{
		users: make(map[UserID]*User),
		port:  port,
		mux:   httprouter.New(),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		publisher:  publisher,
		subscriber: subscriber,
		removeUser: make(chan *User),
	}
	server.initRoute()
	return server
}

func (s *Server) Run() {
	go func() {
		for {
			select {
			case usr := <-s.removeUser:
				log.Println("Removing user")
				log.Println(usr)
				log.Println(len(s.users))
				delete(s.users, usr.UserID)
				log.Println(len(s.users))
			}
		}
	}()

	log.Println("Serving on http://localhost" + s.port)
	log.Println(http.ListenAndServe(s.port, s.mux))
}
