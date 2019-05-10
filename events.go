package gusk

import "encoding/json"

func eventCFG(ctx *Socket) func(dt []byte) {
	return func(dt []byte) {
		resp := new(MessageCFG)
		json.Unmarshal(dt, resp)
		switch resp.Data.Mode {
		case "get":
			ctx.WCFG("set", ctx.ID)
			ctx.prepare <- true
		case "server-closed":
			ctx.FinishForServer <- false
		case "set":
			delete(ctx.Upgrader.us, ctx.ID)
			ctx.ID = resp.Data.Data
			ctx.Upgrader.us[ctx.ID] = ctx
			ctx.Reconection = true
			ctx.prepare <- true
		default:
			ctx.WCFG("message", "El modo no es compatible")
		}

	}
}
