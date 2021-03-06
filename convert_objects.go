package capacitor

import (
	"bytes"
	"encoding/base64"
	"strings"
)

func FluxConnectionResponseToBytes(token string) []byte {
	msg := FLUX_CONNECT + "::{NO-TOPIC}::{NO-PAYLOAD}"
	msg = strings.Replace(msg, "{NO-TOPIC}", "0", 1)
	msg = strings.Replace(msg, "{NO-PAYLOAD}", tobase64(token), 1)
	return []byte(msg)
}

func FluxTopicSubscriptionRequestToBytes(topic string) []byte {
	msg := FLUX_TOPIC_SUBSCRIBE + "::{TOPIC}::{NO-PAYLOAD}"
	msg = strings.Replace(msg, "{TOPIC}", topic, 1)
	msg = strings.Replace(msg, "{NO-PAYLOAD}", tobase64("0"), 1)
	return []byte(msg)
}

func FluxMessageToBytes(flxMsg FluxMessage) []byte {
	msg := FLUX_MESSAGE_TEXT + "::{TOPIC}::{PAYLOAD}"
	msg = strings.Replace(msg, "{TOPIC}", flxMsg.Topic, 1)
	msg = strings.Replace(msg, "{PAYLOAD}", tobase64(string(flxMsg.Payload)), 1)

	return []byte(msg)
}

func bytesToFluxMessage(msgBytes []byte) FluxMessage {
	sections := bytes.Split(msgBytes, []byte("::"))
	if len(sections) != 3 {
		// TODO throw an error or notify downstream somehow
		return FluxMessage{}
	}
	fluxMessage := FluxMessage{
		Topic:   string(sections[1]),
		Payload: frombase64(sections[2]),
	}
	return fluxMessage
}

func bytesToFluxObject(object []byte) (RawFluxObject, bool) {
	sections := bytes.Split(object, []byte("::"))
	if len(sections) != 3 {
		// TODO throw an error or notify downstream somehow
		return RawFluxObject{}, false
	}

	flxO := RawFluxObject{
		_control: sections[0],
		_topic:   sections[1],
		_payload: frombase64(sections[2]),
	}

	return flxO, true
}

func tobase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func frombase64(value []byte) (rv []byte) {
	rv = make([]byte, len(value))
	_, err := base64.StdEncoding.Decode(rv, value)
	if err != nil {
		return []byte{}
	}
	return
}
