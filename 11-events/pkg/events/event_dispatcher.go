package events

import (
	"fmt"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface[any]
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface[any]),
	}
}

func (ed *EventDispatcher) Register(
	eventName string,
	handler EventHandlerInterface[any],
) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return fmt.Errorf(
					"handler for event %s already registered",
					eventName,
				)
			}

		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Unregister(
	eventName string,
	handler EventHandlerInterface[any],
) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(
					ed.handlers[eventName][:i],
					ed.handlers[eventName][i+1:]...,
				)
				return nil
			}
		}
	}

	return fmt.Errorf(
		"handler for event %s not found",
		eventName,
	)
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface[any])
}

func (ed *EventDispatcher) Has(
	eventName string,
	handler EventHandlerInterface[any],
) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface[any]) error {
	if _, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}

		for _, handler := range ed.handlers[event.GetName()] {
			wg.Add(1)
			go handler.Handle(event, wg)
		}

		wg.Wait()

		return nil
	}

	return fmt.Errorf("no handlers for event %s", event.GetName())
}
