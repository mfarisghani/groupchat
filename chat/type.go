package chat

import "github.com/gorilla/websocket"

type SenderID string
type UserID string
type RoomID string

type UserEnterRequest struct {
	RoomID string
	UserID string
	Conn   *websocket.Conn
}

type UserLeaveRequest struct {
	User *User
}
