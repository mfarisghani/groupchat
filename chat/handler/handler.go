package handler

import (
	"log"
	"net/http"

	"github.com/groupchat/chat"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	server   *chat.Server
	port     string
	mux      *httprouter.Router
	upgrader websocket.Upgrader
}

func New(server *chat.Server, port string) *Handler {
	return &Handler{
		server: server,
		port:   port,
		mux:    httprouter.New(),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *Handler) Run() {
	go h.server.Run()

	h.initRoute()
	log.Println("Serving on http://localhost" + h.port)
	log.Println(http.ListenAndServe(h.port, h.mux))
}

func (h *Handler) connectRoom(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	roomID := p.ByName("room_id")
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := uuid.NewV4()

	req := &chat.UserEnterRequest{
		RoomID: roomID,
		UserID: id.String(),
		Conn:   conn,
	}

	h.server.OnUserEnter(req)
}
