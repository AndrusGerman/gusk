package gusk

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// NewSocket create new handler socket
func NewSocket(handler func(*Socket), upgraderConfig ...*Upgrader) func(*gin.Context) {
	// Default Upgrader
	var upgrader *Upgrader
	if len(upgraderConfig) > 0 {
		upgrader = upgraderConfig[0]
	} else {
		upgrader = new(Upgrader)
	}
	if upgrader.us == nil {
		upgrader.us = make(map[string]*Socket)
	}
	// Handler
	return func(ctx *gin.Context) {
		var err error
		// Create USER Model
		socket := &Socket{
			Gin:             ctx,
			Upgrader:        upgrader,
			Connect:         true,
			ID:              createID(),
			closedCicle:     make(chan bool),
			FinishForServer: make(chan bool),
		}
		// Create Conection
		socket.WS, err = upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			println("Error-36: " + socket.ID)
			return
		}
		socket.Conection()
		upgrader.us[socket.ID] = socket
		// OnClose
		defer func() {
			close(socket.closedCicle)
			delete(upgrader.us, socket.ID)
			socket.WCFG("clear-conection", nil)
			defer socket.WS.Close()
		}()
		// Run UserHandler
		handler(socket)
		// Finish
		socket.closedCicle <- true
	}
}

func createID() string {
	return fmt.Sprint(time.Now().Unix())
}
