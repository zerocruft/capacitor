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
	_data        []byte
	_payload     []byte
}

func (fo RawFluxObject) GetType() string {
	return string(fo._type)
}

func (fo RawFluxObject) GetClientToken() string {
	return string(fo._clientToken)
}

func (fo RawFluxObject) GetData() string {
	return string(fo._data)
}

func (fo RawFluxObject) GetPayloadBytes() []byte {
	return fo._payload
}
