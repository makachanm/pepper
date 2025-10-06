package runtime

import (
	"github.com/veandco/go-sdl2/sdl"
)

// EventType defines the type of a GUI event.
type EventType string

const (
	EventTypeQuit            EventType = "quit"
	EventTypeMouseMotion     EventType = "mouse_motion"
	EventTypeMouseButtonDown EventType = "mouse_button_down"
	EventTypeMouseButtonUp   EventType = "mouse_button_up"
	EventTypeKeyDown         EventType = "key_down"
	EventTypeKeyUp           EventType = "key_up"
)

// Event represents a single GUI event.
type Event struct {
	Type    EventType
	X, Y    int
	Button  uint8
	Key     sdl.Keycode
	KeyName string
}

// eventQueue is a thread-safe, channel-based event queue.
type eventQueue struct {
	events chan Event
}

// newEventQueue creates a new event queue.
func newEventQueue() *eventQueue {
	return &eventQueue{
		events: make(chan Event, 1024), // Buffered channel for performance
	}
}

// Enqueue adds an event to the queue.
func (q *eventQueue) Enqueue(e Event) {
	q.events <- e
}

// Dequeue removes and returns the first event from the queue.
// It blocks if the queue is empty.
func (q *eventQueue) Dequeue() Event {
	return <-q.events
}

// EventQueue is the global instance of the event queue.
var EventQueue = newEventQueue()