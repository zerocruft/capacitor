package capacitor

import (
	"github.com/gorilla/websocket"
)

var (
	connections         map[string]*websocket.Conn
	clientWriteChannels map[string]chan FluxMessage
	clientReadChannels  map[string]chan FluxMessage
)

func init() {
	connections = map[string]*websocket.Conn{}
	clientWriteChannels = map[string]chan FluxMessage{}
	clientReadChannels = map[string]chan FluxMessage{}
}
