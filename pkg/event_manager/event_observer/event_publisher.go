package event_observer

import (
	"fmt"
	"github.com/Hesam-Eskandari/go-desktop/pkg/lib/observer"
)

type IEventPublisher[T any] interface {
	AttachSubscriber(subscriber observer.ISubscriber[T]) error
	DetachSubscriberById(id string) error
	NotifySubscribers()
	PushToQueue(T)
}

var idCount = 0

type EventPublisher[T any] struct {
	id int
	observer.IPublisher[T]
	bufferSize            int
	incomingPublisherChan chan T
	isNotifying           bool
	isStopping            chan bool
	isStopped             chan bool
}

func NewEventPublisher[T any](incomingPublisherChan chan T, bufferSize int) *EventPublisher[T] {
	subscriberInitCount := 10
	idCount++
	ip := &EventPublisher[T]{
		id:                    idCount,
		IPublisher:            observer.NewPublisher[T](subscriberInitCount),
		incomingPublisherChan: incomingPublisherChan,
		bufferSize:            bufferSize,
		isStopping:            make(chan bool),
		isStopped:             make(chan bool)}
	return ip
}

func (ip *EventPublisher[T]) AttachSubscriber(subscriber observer.ISubscriber[T]) error {
	return ip.Attach(subscriber)
}

func (ip *EventPublisher[T]) DetachSubscriberById(id string) error {
	return ip.Detach(id)
}

// NotifySubscribers asynchronously updates subscribers with new available input events
// or restart the notifier if this async process was running before
func (ip *EventPublisher[T]) NotifySubscribers() {
	if ip.isNotifying {
		ip.isStopping <- true
		<-ip.isStopped
		ip.isNotifying = false
	}
	go ip.sendEvents()
	ip.isNotifying = true
}

func (ip *EventPublisher[T]) sendEvents() {
	ip.setIncomingPublisherEventChan()
	defer ip.destructIncomingPublisherEventChan()
hookListener:
	for {
		select {
		case <-ip.isStopping:
			break hookListener
		case event := <-ip.incomingPublisherChan:
			ip.Notify(event)
		}
	}
	ip.isStopped <- true
}

func (ip *EventPublisher[T]) setIncomingPublisherEventChan() {
	if ip.incomingPublisherChan != nil {
		return
	}
	ip.incomingPublisherChan = make(chan T, ip.bufferSize)
}

func (ip *EventPublisher[T]) destructIncomingPublisherEventChan() {
	if ip.incomingPublisherChan == nil {
		return
	}
	close(ip.incomingPublisherChan)
}

func (ip *EventPublisher[T]) PushToQueue(event T) {
	fmt.Println("pushing event to queue")
	ip.incomingPublisherChan <- event
}
