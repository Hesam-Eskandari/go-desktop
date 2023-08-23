package mouse_events

import (
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events/coordinate"
)

type MouseEventType string

const (
	MouseButtonUp   MouseEventType = "mouseUp"
	MouseButtonDown MouseEventType = "mouseDown"
	MouseButtonDrag MouseEventType = "mouseDrag"
)

type MouseEvent struct {
	coordinate.EventCoordinate
	mouseEventType MouseEventType
	ButtonCode     uint16
}

func NewMouseEvent(coords coordinate.EventCoordinate, mouseEventType MouseEventType, ButtonCode uint16) MouseEvent {
	return MouseEvent{coords, mouseEventType, ButtonCode}
}

type MouseEventState struct {
	states     chan MouseEvent
	ButtonCode uint16
}

func NewMouseEventState(ButtonCode uint16) MouseEventState {
	states := make(chan MouseEvent, 100)
	return MouseEventState{states, ButtonCode}
}

func (mes *MouseEventState) PushState(mv MouseEvent) {
	mes.states <- mv
}

func (mes *MouseEventState) GetStates() <-chan MouseEvent {
	return mes.states
}

func (mes *MouseEventState) CloseStates() {
	close(mes.states)
}
