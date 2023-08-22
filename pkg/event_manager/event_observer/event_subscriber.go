package event_observer

import (
	"github.com/Hesam-Eskandari/go-desktop/pkg/lib/observer"
)

type eventSubscriber[T any] struct {
	id                  string
	subscriberEventChan chan T
}

type IEventSubscriber[T any] interface {
	observer.ISubscriber[T]
	GetEventChan() <-chan T
}

func NewEventSubscriber[T any](id string, bufferSize int) IEventSubscriber[T] {
	return &eventSubscriber[T]{
		id,
		make(chan T, bufferSize),
	}
}

func (is *eventSubscriber[T]) Update(event T) {
	is.subscriberEventChan <- event
}

func (is *eventSubscriber[T]) GetId() string {
	return is.id
}

func (is *eventSubscriber[T]) GetEventChan() <-chan T {
	return is.subscriberEventChan
}
