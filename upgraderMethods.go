package gusk

import "github.com/gin-gonic/gin"

// WLOG write log to all users
func (ctx *Upgrader) WLOG(data interface{}, notSendUsers ...*Socket) {
	ctx.WJSON("cfg", gin.H{"Mode": "server-log", "Data": data}, notSendUsers...)
}

// WJSON write json to all users
func (ctx *Upgrader) WJSON(event string, data interface{}, notSendUsers ...*Socket) {
	// Send to user?
	send := func(user *Socket) bool {
		for _, val := range notSendUsers {
			if val == user {
				return false
			}
		}
		return true
	}
	// Send WJSON
	for _, val := range ctx.users {
		if send(val) {
			val.WJSON(event, data)
		}
	}
}

// WCFG write message to all users
func (ctx *Upgrader) WCFG(mode string, data interface{}, notSendUsers ...*Socket) {
	ctx.WJSON("cfg", gin.H{"Mode": mode, "Data": data}, notSendUsers...)
}
