package main

import (
	"fmt"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/event_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/input_observer"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events"
	"github.com/Hesam-Eskandari/go-desktop/pkg/event_manager/mouse_events/mouse_button"
	"runtime"
)

type out struct {
	in chan int
}

func newOut() out {
	return out{make(chan int, 5)}
}

func populate(co chan out) {
	for {
		o := newOut()

		for i := 0; i < 3; i++ {
			o.in <- i
			if i == 0 {
				co <- o
			}
		}
		close(o.in)
	}
}

func idea() {
	co := make(chan out, 2)
	go populate(co)
	for o := range co {
		for i := range o.in {
			fmt.Println(i)
		}
	}
}

func main() {
	run()
	//idea()
}

func run() {
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
	clickSub := event_observer.NewEventSubscriber[mouse_events.MouseEventState]("clickSub", 1)
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
		//fmt.Println(event, count)
		for state := range event.GetStates() {
			count++
			fmt.Println(state, count, runtime.NumGoroutine())
		}
	}
}
