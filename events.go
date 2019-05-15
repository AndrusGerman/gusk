package gusk

// Create cfg event
func eventGuskCFG(ctx *Socket) func(interface{}) {
	return func(dt interface{}) {
		resp := dt.(map[string]interface{})
		var mode = resp["Mode"].(string)
		switch mode {
		case ModeServer.GetConfiguration:
			ctx.WCFG(ModeClient.SetConfiguration, ctx.ID)
			ctx.SocketData.CfgPrepareSet(true)
		case ModeServer.CloseGusk:
			ctx.CloseSignal()
		case ModeServer.SetConfigurationReconection:
			// Set Configuration get data
			if dt, ok := resp["Data"].(map[string]interface{}); ok {
				// Set configuration get ID
				if id, ok := dt["ID"].(string); ok {
					// Set configuration set ID
					if ctx.Upgrader.ChangeIDReconection(ctx, id) {
						return
					}
				}
			}
			ctx.WLOGISTRUE(!ctx.Reconection, "Error 2EGC= data no valid") // Sed log on not conection
			ctx.SocketData.CfgPrepareSet(false)
			ctx.CloseSignal()
			// Send is prepare true or false
		default:
			ctx.WLOG("Error 1EGC='" + mode + "' not found")
		}
	}
}
