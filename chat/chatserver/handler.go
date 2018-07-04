package chatserver

import (
	"log"
	"net/http"

	"github.com/groupchat/chat"
	"github.com/julienschmidt/httprouter"
)

func (s *Server) connectRoom(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	roomID := p.ByName("room_id")
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	usr := chat.NewUser(chat.UserID("antony"), chat.RoomID(roomID), "Antony", s.publisher, s.subscriber, conn)
	s.Users[usr.UserID] = usr
	usr.Run()
}
