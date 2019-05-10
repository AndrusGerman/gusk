package gusk

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader model for message
type Upgrader struct {
	websocket.Upgrader
	us map[string]*Socket
}

// Socket model for conection
type Socket struct {
	Gin             *gin.Context
	WS              *websocket.Conn
	ID              string
	Upgrader        *Upgrader
	Connect         bool
	closedCicle     chan bool
	FinishForServer chan bool
	prepare         chan bool
	Reconection     bool
	Event           map[string]func([]byte)
}

// MessageGet model
type MessageGet struct {
	Event string
}

// MessageSend model
type MessageSend struct {
	Event string
	Data  interface{}
}

// MessageCFG configuration
type MessageCFG struct {
	Event string
	Data  struct {
		Mode string
		Data string
	}
}
