package mouse_button

import (
	"fmt"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/event_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/input_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events/coordinate"
	hook "github.com/robotn/gohook"
	"time"
)

const (
	LeftClick  uint16 = 1
	RightClick uint16 = 2
)

type mouseButton struct {
	input_observer.IInputSubscriber
	*event_observer.EventPublisher[mouse_events.MouseEventState]
	buttonCode uint16
}

var countId = 0

func NewMouseButton(buttonCode uint16) IMouseButton {
	countId++
	subscriberId := fmt.Sprintf("mouseButtonSubscriber-%v", countId)
	publisherBufferSize := 1000
	m := &mouseButton{
		input_observer.NewInputSubscriber(subscriberId),
		event_observer.NewEventPublisher[mouse_events.MouseEventState](nil, publisherBufferSize),
		buttonCode,
	}
	//m.IInputSubscriber = m
	go m.connectPubSub()
	return m
}

type IMouseButton interface {
	GetSubscriber() input_observer.IInputSubscriber
	event_observer.IEventPublisher[mouse_events.MouseEventState]
}

func (c *mouseButton) GetSubscriber() input_observer.IInputSubscriber {
	return c.IInputSubscriber
}

func (c *mouseButton) connectPubSub() {
	isMouseDownDetected := false
	var mouseEventState mouse_events.MouseEventState
	var delayStart time.Time = time.Now()
	for ev := range c.GetEventChan() {
		if (ev.Kind == hook.MouseHold || ev.Kind == hook.MouseDown) && !isMouseDownDetected && ev.Button == c.buttonCode {
			//continue
			if time.Since(delayStart) > time.Second*5 || time.Since(delayStart) < time.Millisecond*10 {
				isMouseDownDetected = false
				delayStart = time.Now()
				continue
			}
			isMouseDownDetected = true
			mouseEventState = mouse_events.NewMouseEventState(c.buttonCode)
			coords := coordinate.EventCoordinate{X: ev.X, Y: ev.Y}
			mouseEventState.States <- mouse_events.NewMouseEvent(coords, mouse_events.MouseButtonDown)
			c.PushToQueue(mouseEventState)
		} else if (ev.Kind == hook.MouseUp || ev.Kind == hook.MouseDown) && isMouseDownDetected && ev.Button == c.buttonCode {
			//continue
			delayStart = time.Now()
			coords := coordinate.EventCoordinate{X: ev.X, Y: ev.Y}
			mouseEventState.States <- mouse_events.NewMouseEvent(coords, mouse_events.MouseButtonUp)
			close(mouseEventState.States)
			isMouseDownDetected = false
		}

	}
}
