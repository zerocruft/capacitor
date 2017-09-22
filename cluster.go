package capacitor

type FluxPing struct {
	Node FluxNode `json:"node"`
	NumberOfConnections int `json:"connections"`
}

type FluxPong struct {
	Peers []FluxNode `json:"peers"`
}

type FluxNode struct {
	Address string `json:"address"`
}
