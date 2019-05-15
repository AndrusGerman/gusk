package gusk

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader model for message
type Upgrader struct {
	websocket.Upgrader
	usersSk map[string]*Socket
}

// Socket model for conection
type Socket struct {
	WS *websocket.Conn `json:"-"`
	// ID user conection
	ID   string
	HTTP struct {
		Request *http.Request
	} `json:"-"`
	// Is Socket
	Upgrader    *Upgrader   `json:"-"`
	Connect     bool        `json:"-"`
	Finish      chan error  `json:"-"`
	SocketData  *socketData `json:"-"`
	Reconection bool        `json:"-"`
	OnClose     func()      `json:"-"`
}

type socketData struct {
	events        map[string]func(interface{})
	onClosedError error
	Rooms         []string
	prepare       chan bool
	cfgComplete   bool
}

// Message model
type Message struct {
	Event string
	Data  interface{}
}

// H is a shortcut for map[string]interface{}
type H map[string]interface{}
