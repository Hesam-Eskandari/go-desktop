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

func (es *eventSubscriber[T]) Update(event T) {
	es.subscriberEventChan <- event
}

func (es *eventSubscriber[T]) GetId() string {
	return es.id
}

func (es *eventSubscriber[T]) GetEventChan() <-chan T {
	return es.subscriberEventChan
}
