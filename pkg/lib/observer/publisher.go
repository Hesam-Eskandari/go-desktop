package observer

import "fmt"

type IPublisher[T interface{}] interface {
	Notify(report T)
	Attach(subscriber ISubscriber[T]) error
	Detach(id string) error
}

type publisher[T any] struct {
	subscribers []ISubscriber[T]
}

func NewPublisher[T any](bufferSize int) IPublisher[T] {
	return &publisher[T]{
		subscribers: make([]ISubscriber[T], 0, bufferSize),
	}
}

func (p *publisher[T]) Attach(subscriber ISubscriber[T]) error {
	index := p.findSubscriberIndex(subscriber.GetId())
	if index != -1 {
		return fmt.Errorf("subscriber with id: \"%s\" already exists", subscriber.GetId())
	}
	p.subscribers = append(p.subscribers, subscriber)
	return nil
}

func (p *publisher[T]) Notify(report T) {
	for _, sub := range p.subscribers {
		sub.Update(report)
	}
}

func (p *publisher[T]) Detach(id string) error {
	index := p.findSubscriberIndex(id)
	if index == -1 {
		return fmt.Errorf("error: subscriber with id: \"%s\" does not exist", id)
	}
	p.removeSubscriberByIndex(index)
	return nil
}

func (p *publisher[T]) findSubscriberIndex(id string) int {
	for index, sub := range p.subscribers {
		if sub.GetId() == id {
			return index
		}
	}
	return -1
}

func (p *publisher[T]) removeSubscriberByIndex(index int) {
	p.subscribers = append(p.subscribers[:index], p.subscribers[index+1:]...)
}
