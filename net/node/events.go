package node

import (
	"fmt"

	"github.com/nknorg/nkn/events"
)

type eventQueue struct {
	Consensus  *events.Event
	Block      *events.Event
	Disconnect *events.Event
	Relay      *events.Event
}

func (eq *eventQueue) init() {
	eq.Consensus = events.NewEvent()
	eq.Block = events.NewEvent()
	eq.Disconnect = events.NewEvent()
	eq.Relay = events.NewEvent()
}

func (eq *eventQueue) GetEvent(eventName string) *events.Event {
	switch eventName {
	case "consensus":
		return eq.Consensus
	case "block":
		return eq.Block
	case "disconnect":
		return eq.Disconnect
	case "relay":
		return eq.Relay
	default:
		fmt.Printf("Unknow event")
		return nil
	}
}
