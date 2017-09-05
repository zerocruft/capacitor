package capacitor

import (
	"github.com/gorilla/websocket"
	"sync"
)

var (
	connections      map[string]*websocket.Conn
	connectionsMutex sync.Mutex
)

func init() {
	connections = map[string]*websocket.Conn{}
	connectionsMutex = sync.Mutex{}
}

func addConnection(token string, conn *websocket.Conn) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	connections[token] = conn
}

func removeConnectionAfterClosing(token string) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	connections[token].Close()
	delete(connections, token)
}
