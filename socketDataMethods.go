package gusk

func (ctx *socketData) CfgPrepareSet(bl bool) {
	if ctx.cfgComplete == false {
		ctx.cfgComplete = true
		ctx.prepare <- bl
	}
}
