package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	UserID     UserID `json:"user_id"`
	RoomID     RoomID `json:"room_id"`
	Name       string `json:"name"`
	publisher  Publisher
	subscriber Subscriber
	conn       *websocket.Conn
	message    chan string
}

func NewUser(userID UserID, roomID RoomID, name string, publisher Publisher, subscriber Subscriber, conn *websocket.Conn) *User {
	return &User{
		UserID:     userID,
		RoomID:     roomID,
		Name:       name,
		publisher:  publisher,
		subscriber: subscriber,
		conn:       conn,
	}
}

func (u *User) read() {
	for {
		_, msg, err := u.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			u.conn.Close()
		}

		u.publisher.Publish(u.RoomID, string(msg))
	}
}

func (u *User) write() {
	for {
		select {
		case msg := <-u.message:
			u.conn.WriteMessage(0, []byte(msg))
		}
	}
}

func (u *User) Consume(message string) error {
	u.message <- message
	return nil
}

func (u *User) Run() {
	if err := u.subscriber.Subscribe(u); err != nil {
		log.Println(err)
		return
	}
	go u.read()
	go u.write()
}
