package gusk

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

// /*
// 	Write Methods
// */

// WJSON write json and send
func (ctx *Socket) WJSON(event string, d interface{}) error {
	if ctx.Connect {
		return ctx.WS.WriteJSON(MessageSend{Event: event, Data: d})
	}
	return errors.New("No connect")
}

// WCFG read message and parse
func (ctx *Socket) WCFG(mode string, data interface{}) error {
	return ctx.WJSON("cfg", gin.H{"Mode": mode, "Data": data})
}

// /*
// 	Read Methods
// */

// RJSON read message and parse
func (ctx *Socket) RJSON(d interface{}) error {
	return ctx.WS.ReadJSON(d)
}

// Read read message and parse
func (ctx *Socket) Read() ([]byte, error) {
	_, b, err := ctx.WS.ReadMessage()
	return b, err
}

// Conection init conection
func (ctx *Socket) Conection() {
	var terminar = false
	if ctx.Event == nil {
		ctx.Event = make(map[string]func([]byte))
		ctx.Event["cfg"] = eventCFG(ctx)
	}
	go func() {
		<-ctx.closedCicle
		terminar = true
	}()
	go func() {
		for {
			if terminar {
				return
			}
			resp := new(MessageGet)
			dt, err := ctx.Read()
			if err != nil {
				return
			}
			json.Unmarshal(dt, resp)
			ctx.Event[resp.Event](dt)
		}
	}()
}
