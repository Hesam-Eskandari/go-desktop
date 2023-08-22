package main

import (
	"fmt"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/event_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/input_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events/mouse_button"
)

func main() {
	inputPublisher := input_observer.NewInputPublisher()
	leftMouseButton := mouse_button.NewMouseButton(mouse_button.LeftClick)
	if err := inputPublisher.AttachSubscriber(leftMouseButton.GetSubscriber()); err != nil {
		panic(err)
	}
	rightMouseButton := mouse_button.NewMouseButton(mouse_button.RightClick)
	if err := inputPublisher.AttachSubscriber(rightMouseButton.GetSubscriber()); err != nil {
		panic(err)
	}
	inputPublisher.NotifySubscribers()
	clickSub := event_observer.NewEventSubscriber[mouse_events.MouseEventState]("clickSub", 100)
	if err := leftMouseButton.AttachSubscriber(clickSub); err != nil {
		panic(err)
	}
	if err := rightMouseButton.AttachSubscriber(clickSub); err != nil {
		panic(err)
	}
	leftMouseButton.NotifySubscribers()
	rightMouseButton.NotifySubscribers()

	printEvents(clickSub.GetEventChan())
}

func printEvents(c <-chan mouse_events.MouseEventState) {
	for event := range c {
		count := 0
		for state := range event.States {
			count++
			fmt.Println(state, count)
		}
	}
}
