package capacitor

type FluxPing struct {
	Node                FluxNode `json:"node"`
	NumberOfConnections int      `json:"connections"`
}

type FluxPong struct {
	Peers []FluxNode `json:"peers"`
}

type FluxNode struct {
	ClientEndpoint string `json:"endpoint"`
	Name           string `json:"name"`
	PeerEndpoint   string `json:"peer-endpoint"`
}
