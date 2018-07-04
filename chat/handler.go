package chat

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) connectRoom(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	roomID := p.ByName("room_id")
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	usr := NewUser(s, UserID("antony"), RoomID(roomID), "Antony", s.publisher, s.subscriber, conn)
	s.users[usr.UserID] = usr
	usr.Run()

	log.Println("New user connected", usr)
	log.Println(len(s.users))
}
