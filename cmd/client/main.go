package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

// const address string = "172.31.5.228:8080"

const address string = "localhost:8080"

func main() {

	initWebsocketClient()
}

func initWebsocketClient() {
	fmt.Println("Starting Client")
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/room/haha", address), "", fmt.Sprintf("http://%s/", address))
	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}
	// incomingMessages := make(chan string)
	// go readClientMessages(ws, incomingMessages)
	// i := 0
	for i := 0; i <= 1000000; i++ {
		// select {
		// case <-time.After(time.Duration(2e9)):
		// i++
		// response := new(Message)
		// response.RequestID = i
		// response.Command = "Eject the hot dog."
		// err = websocket.JSON.Send(ws, response)
		go func() {
			_ = websocket.Message.Send(ws, fmt.Sprintf(`{"id":"553943b0-7f69-11e8-80fd-c78bd0801asd","username":"user_%v","message":"goks"}`, i%50))
		}()
		// if err != nil {
		// 	fmt.Printf("Send failed: %s\n", err.Error())
		// 	// os.Exit(1)
		// }
		time.Sleep(500 * time.Millisecond)
		log.Println("sent!")
		// case message := <-incomingMessages:
		// 	fmt.Println(`Message Received:`, message)

		// }
	}
}

func readClientMessages(ws *websocket.Conn, incomingMessages chan string) {
	for {
		var message string
		// err := websocket.JSON.Receive(ws, &message)
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Printf("Error::: %s\n", err.Error())
			return
		}
		incomingMessages <- message
	}
}
