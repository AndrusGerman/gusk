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
			println("Error-37: " + socket.ID)
			return
		}
		socket.prepare = make(chan bool)
		socket.Conection()
		upgrader.us[socket.ID] = socket
		// OnClosed conection
		socket.WS.SetCloseHandler(func(a int, b string) error {
			socket.Connect = false
			socket.FinishForServer <- true
			return nil
		})
		// onClosed Handler
		defer func() {
			close(socket.closedCicle)
			close(socket.FinishForServer)
			close(socket.prepare)
			delete(upgrader.us, socket.ID)
			if socket.Connect {
				socket.WCFG("clear-conection", nil)
			}
			socket.WS.Close()
		}()

		// Run UserHandler
		<-socket.prepare
		handler(socket)
		// Finish
		socket.closedCicle <- true
	}
}

func createID() string {
	return fmt.Sprint(time.Now().Unix())
}
