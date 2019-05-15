package gusk

import (
	"errors"
)

/*
	Write Methods
*/

// Send write messaga and send
func (ctx *Socket) Send(event string, d interface{}) error {
	if ctx.Connect {
		return ctx.WS.WriteJSON(Message{Event: event, Data: d})
	}
	return errors.New("Error 1SM=Not conection")
}

// WCFG read message and parse
func (ctx *Socket) WCFG(mode string, data interface{}) error {
	return ctx.Send("cfg", H{"Mode": mode, "Data": data})
}

// WLOG read message and send
func (ctx *Socket) WLOG(data interface{}) error {
	return ctx.WCFG(ModeClient.Log, data)
}

// WLOGISTRUE read message and send if true
func (ctx *Socket) WLOGISTRUE(boolVar bool, data interface{}) bool {
	if boolVar {
		ctx.WLOG(data)
	}
	return boolVar
}

// CloseSignal websocket connection
func (ctx *Socket) CloseSignal() error {
	ctx.SocketData.onClosedError = ctx.WCFG(ModeClient.CloseGusk, nil)
	return ctx.SocketData.onClosedError
}

// ForceClose websocket disconnection
func (ctx *Socket) ForceClose() {
	ctx.Finish <- nil
}

/*
	Read Methods
*/

// ReadMessage read message and parse
func (ctx *Socket) ReadMessage() (*Message, error) {
	var mensaje = new(Message)
	return mensaje, ctx.WS.ReadJSON(mensaje)
}

// Read read message and parse
func (ctx *Socket) Read() ([]byte, error) {
	_, b, err := ctx.WS.ReadMessage()
	return b, err
}

// AddRoom for user
func (ctx *Socket) AddRoom(room string) {
	ctx.SocketData.Rooms = append(ctx.SocketData.Rooms, room)
}

// CicleEvents init conection
func (ctx *Socket) CicleEvents() {
	ctx.SocketData.events = make(map[string]func(interface{}))
	// Set Configuration Route
	ctx.Event("cfg", eventGuskCFG(ctx))
	// Conection Cilcle
	go func() {
		for {
			// Parse Response
			resp, err := ctx.ReadMessage()
			// Check error
			if err != nil {
				ctx.WLOG("Error 2SM=" + err.Error())
				return
			}
			// Check Event
			if ctx.SocketData.events[resp.Event] == nil {
				ctx.WLOG("Error 3SM=Event '" + resp.Event + "' not found ")
			} else {
				// Send Message
				ctx.SocketData.events[resp.Event](resp.Data)
			}
			// Check conection
			if ctx.Connect == false {
				return
			}
		}
	}()
}

// Event read message and parse
func (ctx *Socket) Event(eventName string, handler func(interface{})) {
	ctx.SocketData.events[eventName] = handler
}
