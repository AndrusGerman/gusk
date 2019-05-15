package gusk

import "errors"

// WLOG write log to all users
func (ctx *Upgrader) WLOG(data interface{}, room string) {
	ctx.WCFG(ModeClient.Log, data, room)
}

// Send write json to all users
func (ctx *Upgrader) Send(event string, data interface{}, room string) {
	// Send WJSON
	for _, val := range ctx.UsersInRoom(room) {
		val.Send(event, data)
	}
}

// SendArray write json to all users
func (ctx *Upgrader) SendArray(room string, message []Message) []error {
	var errores []error
	// Send WJSON
	for _, val := range ctx.UsersInRoom(room) {
		for ind := range message {
			errores = append(errores, val.Send(message[ind].Event, message[ind].Data))
		}
	}
	return errores
}

// SendMasive write json to all users
func (ctx *Upgrader) SendMasive(event string, data interface{}, notSend ...*Socket) {
	// Send WJSON
	for _, val := range ctx.getUsers(notSend...) {
		val.Send(event, data)
	}
}

// SendMasive write json to all users
func (ctx *Upgrader) getUsers(notReturn ...*Socket) []*Socket {
	var users []*Socket
	//
	contain := func(user *Socket) bool {
		for indc := range notReturn {
			if user == notReturn[indc] {
				return true
			}
		}
		return false
	}

	//
	for _, val := range ctx.usersSk {
		if !contain(val) {
			users = append(users, val)
		}
	}
	return users
}

// UsersInRoom get users in room
func (ctx *Upgrader) UsersInRoom(room string) []*Socket {
	var users []*Socket
	for indUsers := range ctx.usersSk {
		for indRoom := range ctx.usersSk[indUsers].SocketData.Rooms {
			if ctx.usersSk[indUsers].SocketData.Rooms[indRoom] == room {
				users = append(users, ctx.usersSk[indUsers])
				break
			}
		}
	}
	return users
}

// WCFG write message to all users
func (ctx *Upgrader) WCFG(mode string, data interface{}, room string) {
	ctx.Send("cfg", H{"Mode": mode, "Data": data}, room)
}

// ChangeID to user
func (ctx *Upgrader) ChangeID(user *Socket, newID string) (change bool) {
	if len(newID) == 0 {
		return false
	}
	if user.ID != "" {
		if ctx.usersSk[user.ID] != nil {
			delete(ctx.usersSk, user.ID)
		}
	}
	user.ID = newID
	ctx.usersSk[newID] = user
	return true
}

// ChangeID to user
func (ctx *Upgrader) ChangeIDReconection(user *Socket, newID string) (change bool) {
	if ctx.ChangeID(user, newID) {
		user.Reconection = true // Reconection set
		user.SocketData.CfgPrepareSet(true)
		return true
	}
	return false
}

// Users get connect users
func (ctx *Upgrader) Users() []*Socket {
	var users []*Socket
	for indUsers := range ctx.usersSk {
		if ctx.usersSk[indUsers].Connect {
			users = append(users, ctx.usersSk[indUsers])
		}
	}
	return users
}

// GetUser by ID
func (ctx *Upgrader) GetUser(ID string) (err error, user *Socket) {
	user = ctx.usersSk[ID]
	if user == nil {
		return errors.New("User is nil"), user
	}
	if user.Connect == false {
		return errors.New("User no connect"), user
	}
	return nil, user
}
