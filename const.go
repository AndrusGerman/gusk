package gusk

var (
	// ModeClient Modes cfg: for server send to client
	ModeClient = struct {
		Log              string
		SetConfiguration string
		CloseGusk        string
	}{
		Log:              "server->client:log",
		SetConfiguration: "server->client:set-configuration",
		CloseGusk:        "server->client:close-gusk",
	}
	// ModeServer Modes cfg: for client send to server
	ModeServer = struct {
		CloseGusk                   string
		GetConfiguration            string
		SetConfigurationReconection string
	}{
		CloseGusk:                   "client->server:close-gusk",
		GetConfiguration:            "client->server:get-configuration",
		SetConfigurationReconection: "client->server:set-configuration-reconection",
	}
)
