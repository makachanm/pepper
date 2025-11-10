package gfx

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

// DequeueNonBlocking tries to dequeue an event without blocking.
// It returns the event and true if successful, or a zero-value event and false if the queue is empty.
func (q *eventQueue) DequeueNonBlocking() (Event, bool) {
	select {
	case event := <-q.events:
		return event, true
	default:
		return Event{}, false
	}
}

func (q *eventQueue) IsEmpty() bool {
	return len(q.events) == 0
}

// EventQueue is the global instance of the event queue.
var EventQueue = newEventQueue()

func (pg *PepperGraphics) runEventLoop() {
	defer pg.wg.Done()
	for !ShouldQuit {
		event := sdl.PollEvent()
		if event != nil {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				EventQueue.Enqueue(Event{Type: EventTypeQuit})
				ShouldQuit = true // Signal VM to quit
			case *sdl.MouseMotionEvent:
				EventQueue.Enqueue(Event{
					Type: EventTypeMouseMotion,
					X:    int(e.X),
					Y:    int(e.Y),
				})
			case *sdl.MouseButtonEvent:
				var eventType EventType
				if e.State == sdl.PRESSED {
					eventType = EventTypeMouseButtonDown
				} else {
					eventType = EventTypeMouseButtonUp
				}
				EventQueue.Enqueue(Event{
					Type:   eventType,
					X:      int(e.X),
					Y:      int(e.Y),
					Button: e.Button,
				})
			case *sdl.KeyboardEvent:
				var eventType EventType
				if e.State == sdl.PRESSED {
					eventType = EventTypeKeyDown
				} else {
					eventType = EventTypeKeyUp
				}
				EventQueue.Enqueue(Event{
					Type:    eventType,
					Key:     e.Keysym.Sym,
					KeyName: sdl.GetKeyName(e.Keysym.Sym),
				})
			}
		}
		// Small delay to prevent busy-waiting
	}
	pg.Surface.Finish()
	pg.Window.Destroy()
	sdl.CloseAudio()
	sdl.Quit()
}
