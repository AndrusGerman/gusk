package gusk

import (
	"fmt"
	"net/http"
)

// NewSocket create new handler socket
func NewSocket(handler func(*Socket), upgraderConfig ...*Upgrader) func(http.ResponseWriter, *http.Request) {
	// Default Upgrader
	var upgrader *Upgrader
	// Check upgrader
	if len(upgraderConfig) > 0 {
		upgrader = upgraderConfig[0]
	} else {
		upgrader = new(Upgrader)
	}
	// Check users in upgrader
	if upgrader.users == nil {
		upgrader.users = make(map[string]*Socket)
	}
	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		// Create Model
		socket := &Socket{
			Upgrader: upgrader,
			Connect:  true,
		}
		// Set Request
		socket.HTTP.Request = r
		// Create Conection
		socket.WS, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		// Add Chanels
		socket.FinishForServer = make(chan bool)
		socket.prepare = make(chan bool)
		// Add UserData
		socket.ID = randStringBytesMaskImprSrcUnsafe(14)
		upgrader.users[socket.ID] = socket
		// Add Conection Events
		socket.CicleEvents()
		// OnClosed conection
		socket.WS.SetCloseHandler(func(a int, b string) error {
			// Finish emit
			socket.FinishForServer <- true
			// Conect is false
			socket.Connect = false
			// Send False on error not init handler
			socket.prepare <- false
			// Clear chanels
			close(socket.prepare)
			return nil
		})
		// defer close function
		defer func() {
			fmt.Println("Defer Handler")
			// Clear Connection
			if socket.Connect {
				socket.WCFG("close-configuration", nil)
				socket.WS.Close()
			}
			// Clear vars
			delete(upgrader.users, socket.ID)
			close(socket.FinishForServer)
		}()

		// Run UserHandler
		if <-socket.prepare {
			handler(socket)
		}
	}
}
