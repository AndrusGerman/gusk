package gusk

import (
	"errors"
	"net/http"
)

func (socket *Socket) setInitDataSocket(upgrader *Upgrader) {
	// Add UserData
	socket.ID = randStringBytesMaskImprSrcUnsafe(14)
	socket.SocketData = &socketData{
		onClosedError: errors.New("Error closed conection"),
		prepare:       make(chan bool, 2),
		Rooms:         []string{socket.ID},
	}
	socket.OnClose = func() {

	}
	// Add Chanels
	socket.Finish = make(chan error, 2)
	// Add Upgrader
	socket.Upgrader = upgrader
	// socket room
}

func setInitSocketConnection(w http.ResponseWriter, r *http.Request, upgrader *Upgrader) (*Socket, error) {
	var err error
	var socket = new(Socket)
	socket.WS, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	// Set Vars
	socket.Connect = true
	socket.HTTP.Request = r
	return socket, nil
}

func setInitUpgrader(up ...*Upgrader) *Upgrader {
	// Default Upgrader
	var upgrader *Upgrader
	// Check upgrader
	if len(up) > 0 {
		upgrader = up[0]
	} else {
		upgrader = new(Upgrader)
	}
	// Check users in upgrader
	if upgrader.usersSk == nil {
		upgrader.usersSk = make(map[string]*Socket)
	}
	return upgrader
}

func setCloseGuskUser(upg *Upgrader, socket *Socket) {
	// Clear Connection
	if socket.Connect {
		socket.WCFG("server->client:close-gusk", nil)
		socket.WS.Close()
	}
	// Clear vars
	delete(upg.usersSk, socket.ID)
	close(socket.Finish)
}
