package core

type Event struct {
    Type    string
    Payload interface{}
}

var eventBus = make(chan Event, 100)