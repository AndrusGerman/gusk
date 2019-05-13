package gusk

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
	Write Methods
*/

// WJSON write json and send
func (ctx *Socket) WJSON(event string, d interface{}) error {
	if ctx.Connect {
		return ctx.WS.WriteJSON(Message{Event: event, Data: d})
	}
	return errors.New("Error 1SM=Not conection")
}

// WCFG read message and parse
func (ctx *Socket) WCFG(mode string, data interface{}) error {
	return ctx.WJSON("cfg", gin.H{"Mode": mode, "Data": data})
}

// WLOG read message and parse
func (ctx *Socket) WLOG(data interface{}) error {
	return ctx.WJSON("cfg", gin.H{"Mode": "server-log", "Data": data})
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

// CicleEvents init conection
func (ctx *Socket) CicleEvents() {
	// Set Configuration Route
	ctx.Event = make(map[string]func(interface{}))
	ctx.Event["cfg"] = eventGuskCFG(ctx)
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
			if ctx.Event[resp.Event] == nil {
				ctx.WLOG(fmt.Sprintf("Error 3SM=Event '%s' not found ", resp.Event))
			} else {
				// Send Message
				ctx.Event[resp.Event](resp.Data)
			}
			// Check conection
			if ctx.Connect == false {
				return
			}
		}
	}()
}
