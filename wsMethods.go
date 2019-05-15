package gusk

func wsCloseHandler(socket *Socket) func(int, string) error {
	return func(a int, b string) error {
		// Finish emit
		socket.Finish <- socket.SocketData.onClosedError
		// Conect is false
		socket.Connect = false
		// Send False on error not init handler
		socket.SocketData.CfgPrepareSet(false)
		// Clear chanels
		close(socket.SocketData.prepare)
		socket.OnClose()
		return nil
	}
}
