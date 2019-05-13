package gusk

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader model for message
type Upgrader struct {
	websocket.Upgrader
	users map[string]*Socket
}

// Socket model for conection
type Socket struct {
	WS *websocket.Conn
	// ID user conection
	ID   string
	HTTP struct {
		Request *http.Request
	}
	// Is Socket
	Upgrader        *Upgrader
	Connect         bool
	FinishForServer chan bool
	prepare         chan bool
	Reconection     bool
	Event           map[string]func(interface{})
}

// Message model
type Message struct {
	Event string
	Data  interface{}
}
