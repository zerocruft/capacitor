package capacitor

type FluxMessage struct {
	Control string
	Topic   string
	Payload []byte
}

func (fm FluxMessage) IsZero() bool {

	//TODO this might be deprecated
	return false
}

type fluxMessageTopicSubscribe struct {
}

type RawFluxObject struct {
	_control []byte
	_topic   []byte
	_payload []byte
}

func (fo RawFluxObject) GetType() string {
	return string(fo._control)
}

func (fo RawFluxObject) GetTopic() string {
	return string(fo._topic)
}

func (fo RawFluxObject) GetPayloadBytes() []byte {
	return fo._payload
}
