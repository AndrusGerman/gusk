package gusk

// Create cfg event
func eventGuskCFG(ctx *Socket) func(interface{}) {
	return func(dt interface{}) {
		resp := dt.(map[string]interface{})
		var mode = resp["Mode"].(string)
		switch resp["Mode"] {
		case "get-configuration-server":
			ctx.WCFG("set-configuration", ctx.ID)
			ctx.prepare <- true
		case "close-server":
			ctx.FinishForServer <- false
		case "set-configuration-server":
			delete(ctx.Upgrader.users, ctx.ID)
			ctx.ID = resp["Data"].(string)
			ctx.Upgrader.users[ctx.ID] = ctx
			ctx.Reconection = true
			ctx.prepare <- true
		default:
			ctx.WLOG("Error 1EGC='" + mode + "' not found")
		}
	}
}
