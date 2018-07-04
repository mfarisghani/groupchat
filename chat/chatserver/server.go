package chatserver

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/groupchat/chat"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

type Server struct {
	Users      map[chat.UserID]*chat.User
	port       string
	mux        *httprouter.Router
	upgrader   websocket.Upgrader
	publisher  chat.Publisher
	subscriber chat.Subscriber
}

func New(port string, publisher chat.Publisher, subscriber chat.Subscriber) *Server {
	server := &Server{
		Users: make(map[chat.UserID]*chat.User),
		port:  port,
		mux:   httprouter.New(),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		publisher:  publisher,
		subscriber: subscriber,
	}
	server.initRoute()
	return server
}

func (s *Server) Run() {
	log.Println("Serving on http://localhost" + s.port)
	log.Println(http.ListenAndServe(s.port, s.mux))
}
