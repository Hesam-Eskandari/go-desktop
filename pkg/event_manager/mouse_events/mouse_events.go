package mouse_events

import (
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events/coordinate"
)

type MouseEventType string

const (
	MouseButtonUp   MouseEventType = "mouseUp"
	MouseButtonDown MouseEventType = "mouseDown"
)

type MouseEvent struct {
	coordinate.EventCoordinate
	mouseEventType MouseEventType
}

func NewMouseEvent(coords coordinate.EventCoordinate, mouseEventType MouseEventType) MouseEvent {
	return MouseEvent{coords, mouseEventType}
}

type MouseEventState struct {
	States     chan MouseEvent
	ButtonCode uint16
}

func NewMouseEventState(ButtonCode uint16) MouseEventState {
	return MouseEventState{make(chan MouseEvent, 2), ButtonCode}
}
