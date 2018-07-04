package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	server     *Server
	UserID     UserID `json:"user_id"`
	RoomID     RoomID `json:"room_id"`
	Name       string `json:"name"`
	publisher  Publisher
	subscriber Subscriber
	conn       *websocket.Conn
	message    chan string
	readClose  chan bool
	writeClose chan bool
}

func NewUser(server *Server, userID UserID, roomID RoomID, name string, publisher Publisher, subscriber Subscriber, conn *websocket.Conn) *User {
	return &User{
		server:     server,
		UserID:     userID,
		RoomID:     roomID,
		Name:       name,
		publisher:  publisher,
		subscriber: subscriber,
		conn:       conn,
		message:    make(chan string),
		readClose:  make(chan bool),
		writeClose: make(chan bool),
	}
}

func (u *User) read() {
	for {
		select {
		case isReadClose := <-u.readClose:
			if isReadClose {
				u.server.removeUser <- u
				break
			}
		default:
			_, msg, err := u.conn.ReadMessage()
			if err != nil {
				log.Println(err)
				u.conn.Close()
				u.close <- true
				break
			}

			u.publisher.Publish(u.RoomID, string(msg))
		}
	}
}

func (u *User) write() {
	for {
		select {
		case msg := <-u.message:
			if err := u.conn.WriteMessage(1, []byte(msg)); err != nil {
				log.Println(err)
				u.conn.Close()
				u.close <- true
				break
			}
		case isClose := <-u.close:
			if isClose {
				u.server.removeUser <- u
				break
			}
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
