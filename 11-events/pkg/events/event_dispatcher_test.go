package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface[any], wg *sync.WaitGroup) {
	wg.Done()
}

type MockedHandler struct {
	mock.Mock
}

func (m *MockedHandler) Handle(event EventInterface[any], wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()

	s.handler = TestEventHandler{ID: 1}
	s.handler2 = TestEventHandler{ID: 2}

	s.event = TestEvent{
		Name:    "event 1",
		Payload: "event 1 payload",
	}
	s.event2 = TestEvent{
		Name:    "event 2",
		Payload: "event 2 payload",
	}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Register the first handler
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Register the second handler
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	// Register the first handler
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Register the first handler again
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.Error(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Unregister() {
	// Register the first handler to event
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Register the second handler to event
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Unregister the first handler
	err = s.eventDispatcher.Unregister(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Check if the first handler is not registered
	s.False(s.eventDispatcher.Has(s.event.GetName(), &s.handler))

	// Check if the second handler is registered
	s.True(s.eventDispatcher.Has(s.event.GetName(), &s.handler2))

	// Unregister the second handler
	err = s.eventDispatcher.Unregister(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Check if the second handler is not registered
	s.False(s.eventDispatcher.Has(s.event.GetName(), &s.handler2))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Register the first handler to event
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Register the second handler to event
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// Register the first handler to event 2
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handler)
	s.NoError(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	// Register the second handler to event 2
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handler2)
	s.NoError(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	// Clear the event dispatcher
	s.eventDispatcher.Clear()
	s.Equal(0, len(s.eventDispatcher.handlers[s.event.GetName()]))
	s.Equal(0, len(s.eventDispatcher.handlers[s.event2.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Register the first handler to event
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handler)
	s.NoError(err)

	// Check if the handler is registered
	s.True(s.eventDispatcher.Has(s.event.GetName(), &s.handler))

	// Check if the handler is not registered
	s.False(s.eventDispatcher.Has(s.event.GetName(), &s.handler2))

}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	// Create mocked handler
	mockedHandler := &MockedHandler{}
	mockedHandler.On("Handle", &s.event)

	// Register the mocked handler
	s.eventDispatcher.Register(s.event.GetName(), mockedHandler)

	// Dispatch the event
	err := s.eventDispatcher.Dispatch(&s.event)
	s.NoError(err)

	// Assert that the mocked handler was called
	mockedHandler.AssertCalled(s.T(), "Handle", &s.event)

	// Dispatch event with no handlers registered
	err = s.eventDispatcher.Dispatch(&s.event2)
	s.Error(err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
