package events

import (
	"sync"
	"time"
)

type EventInterface[T any] interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() T
}

type EventHandlerInterface[T any] interface {
	Handle(event EventInterface[T], wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface[any]) error
	Unregister(eventName string, handler EventHandlerInterface[any]) error
	Has(eventName string, handler EventHandlerInterface[any]) bool
	Dispatch(event EventInterface[any]) error
	Clear()
}
