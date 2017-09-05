package capacitor

type FluxMessage struct {
	Topic   string
	Payload []byte
}

func (fm FluxMessage) IsZero() bool {

	//TODO this might be deprecated
	return false
}

type RawFluxObject struct {
	_clientToken []byte
	_type        []byte
	_topic       []byte
	_payload     []byte
}

func (fo RawFluxObject) GetType() string {
	return string(fo._type)
}

func (fo RawFluxObject) GetClientToken() string {
	return string(fo._clientToken)
}

func (fo RawFluxObject) GetTopic() string {
	return string(fo._topic)
}

func (fo RawFluxObject) GetPayloadBytes() []byte {
	return fo._payload
}
