package input_observer

import (
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/event_observer"
	hook "github.com/robotn/gohook"
)

type inputSubscriber struct {
	event_observer.IEventSubscriber[hook.Event]
}

type IInputSubscriber interface {
	event_observer.IEventSubscriber[hook.Event]
}

func NewInputSubscriber(id string) IInputSubscriber {
	inputCacheSize := 10
	return &inputSubscriber{event_observer.NewEventSubscriber[hook.Event](id, inputCacheSize)}
}
