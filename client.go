package capacitor

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

// Returns FluxClient with a write and read channel
func NewClient(req FluxConnectParameters) (FluxClient, error) {

	fluxUrl := "ws://" + req.FluxAddress + ":" + strconv.Itoa(req.FluxPort) + "/flux"
	conn, _, err := websocket.DefaultDialer.Dial(fluxUrl, nil)
	if err != nil {
		fmt.Println(err)
		return FluxClient{}, err
	}

	mt, payload, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		log.Println(mt)
		log.Println(err)
		return FluxClient{}, err
	}

	fluxServiceResponse, success := bytesToFluxObject(payload)
	if !success {
		conn.Close()
		return FluxClient{}, errors.New("Internal Flux Object Error. Invalid msg format")
	}

	//Check to make sure flux object is of connection response type
	if fluxServiceResponse.GetType() != FLUX_CONNECT {
		conn.Close()
		return FluxClient{}, errors.New("Flux Service failed to respond")
	}

	var clientToken = string(fluxServiceResponse.GetPayloadBytes())
	//Now.. subscribe to topics
	for _, topics := range req.Topics {
		channelSubscribeReq := FluxTopicSubscriptionRequestToBytes(topics)
		conn.WriteMessage(websocket.TextMessage, channelSubscribeReq)
	}

	clientWriteChannel := make(chan FluxMessage, 25)
	clientReadChannel := make(chan FluxMessage, 25)
	deadConnectionChannel := make(chan bool, 1)
	addConnection(clientToken, conn)

	// Write Channel
	go func() {
		for {
			newMsg := <-clientWriteChannel
			flxMsgBytes := FluxMessageToBytes(newMsg)
			err := conn.WriteMessage(websocket.TextMessage, flxMsgBytes)
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}
		}
	}()

	// Read Channel
	go func() {
		for {
			mt, payload, err := conn.ReadMessage()
			if mt != websocket.TextMessage || err != nil {
				log.Println(mt)
				log.Println(err)
				conn.Close()
				deadConnectionChannel <- true
				return
			}

			clientReadChannel <- bytesToFluxMessage(payload)
		}
	}()

	return FluxClient{
		clientToken: clientToken,
		send:        &clientWriteChannel,
		receive:     &clientReadChannel,
		token:       clientToken,
		dead:        &deadConnectionChannel,
	}, nil
}

type FluxClient struct {
	clientToken string
	send        *chan FluxMessage
	receive     *chan FluxMessage
	token       string
	closed      bool
	mutex       sync.Mutex
	dead        *chan bool
}

func (fc FluxClient) Token() string {
	return fc.token
}

func (fc *FluxClient) Send() chan FluxMessage {
	return *fc.send
}

func (fc *FluxClient) Receive() chan FluxMessage {
	return *fc.receive
}

func (fc *FluxClient) Dead() chan bool {
	return *fc.dead
}

func (fc FluxClient) AddTopic(topic string) {
	connections[fc.token].WriteMessage(websocket.TextMessage, FluxTopicSubscriptionRequestToBytes(topic))
}

func (fc *FluxClient) Close() {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()

	if fc.closed {
		return
	}

	removeConnectionAfterClosing(fc.token)
	fc.closed = true
}
