module SystemgeOneNodePing

go 1.22.3

//replace github.com/neutralusername/Systemge => ../Systemge
require (
	github.com/gorilla/websocket v1.5.3
	github.com/neutralusername/Systemge v0.0.0-20240804135245-e5af582e4658
)

require golang.org/x/oauth2 v0.21.0 // indirect
