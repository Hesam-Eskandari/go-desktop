package input_observer

import (
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/event_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/lib/observer"
	hook "github.com/robotn/gohook"
)

type IInputPublisher interface {
	AttachSubscriber(subscriber observer.ISubscriber[hook.Event]) error
	DetachSubscriberById(string) error
	NotifySubscribers()
}

type inputPublisher struct {
	*event_observer.EventPublisher[hook.Event]
}

var singletonInputPublisher *inputPublisher = nil

func NewInputPublisher() IInputPublisher {
	if singletonInputPublisher != nil {
		return singletonInputPublisher
	}
	bufferSize := 1000
	singletonInputPublisher = &inputPublisher{event_observer.NewEventPublisher[hook.Event](hook.Start(), bufferSize)}
	return singletonInputPublisher
}

func (ip *inputPublisher) destructIncomingPublisherEventChan() {
	hook.End()
}
