package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	server        *Server
	UserID        UserID `json:"user_id"`
	RoomID        RoomID `json:"room_id"`
	Name          string `json:"name"`
	publisher     Publisher
	subscriber    Subscriber
	conn          *websocket.Conn
	message       chan string
	isReadClosed  bool
	isWriteClosed bool
	readClose     chan bool
	writeClose    chan bool
	checkClose    chan bool
	userClose     chan bool
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
		checkClose: make(chan bool),
		userClose:  make(chan bool),
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
	go u.check()

	<-u.userClose
	log.Println("User need to be closed")

	u.close()
}

func (u *User) read() {
	for {
		select {
		case <-u.readClose:
			log.Println("Read close dipanggil dari channel")
			goto close
		default:
			_, msg, err := u.conn.ReadMessage()
			if err != nil {
				log.Println(err)
				log.Println("Read close dipanggil dari error")
				goto close
			}
			log.Println("Reading", string(msg))
			u.publisher.Publish(u.RoomID, string(msg))
		}
	}
close:
	if !u.isWriteClosed {
		u.writeClose <- true
	}
	u.isReadClosed = true
	u.checkClose <- true
}

func (u *User) write() {
	for {
		select {
		case <-u.writeClose:
			log.Println("Write close dipanggil dari channel")
			goto close
		case msg := <-u.message:
			log.Println("Writing", string(msg))
			if err := u.conn.WriteMessage(1, []byte(msg)); err != nil {
				log.Println(err)
				log.Println("Write close dipanggil dari error")
				goto close
			}
		}
	}
close:
	if !u.isReadClosed {
		u.readClose <- true
	}
	u.isWriteClosed = true
	u.checkClose <- true
}

func (u *User) check() {
	for {
		select {
		case <-u.checkClose:
			log.Println("check close dipanggil", u.isReadClosed, u.isWriteClosed)
			if u.isReadClosed && u.isWriteClosed {
				u.userClose <- true
				break
			}
		}
	}
}

func (u *User) close() {
	log.Println("starting close")

	u.conn.Close()
	u.server.userLeave <- &UserLeaveRequest{
		User: u,
	}
}
