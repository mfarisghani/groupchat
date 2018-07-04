package chat

type Publisher interface {
	Publish(roomID RoomID, message string) error
}

type Subscriber interface {
	Subscribe(user *User) error
}

type Consumer interface {
	Consume(message string) error
}
