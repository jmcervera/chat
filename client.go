package main

import "github.com/gorilla/websocket"

type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is achannel on which messages are sent.
	send chan []byte
	// room is the room this client is chatting on
	room *room
}

// read from the socket via ReadMessage
// sending any received messages to the forward channel on the room type
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// continually accepts messages from the send channel
// writing everything out of the socket via WriteMessage method
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
