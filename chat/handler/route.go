package handler

func (h *Handler) initRoute() {
	h.mux.GET("/room/:room_id", h.connectRoom)
}
