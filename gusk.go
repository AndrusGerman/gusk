package gusk

import (
	"net/http"
)

// NewSocket create new handler socket
func NewSocket(handler func(*Socket), upgraderConfig ...*Upgrader) func(http.ResponseWriter, *http.Request) {
	// init upgrader var
	var upgrader = setInitUpgrader(upgraderConfig...)
	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Create Model
		var socket, err = setInitSocketConnection(w, r, upgrader)
		if err != nil {
			return
		}
		// Set Data
		socket.setInitDataSocket(upgrader)
		upgrader.usersSk[socket.ID] = socket
		// Add Conection Events init clicle events
		socket.CicleEvents()
		// defer close gusk user
		defer setCloseGuskUser(upgrader, socket)
		// OnClosed conection
		socket.WS.SetCloseHandler(wsCloseHandler(socket))
		// Run UserHandler
		if <-socket.SocketData.prepare {
			handler(socket)
		}
	}
}
