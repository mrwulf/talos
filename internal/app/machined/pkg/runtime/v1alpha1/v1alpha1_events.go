// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha1

import (
	"fmt"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"

	"github.com/talos-systems/talos/api/machine"
	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
)

// Events represents the runtime event stream.
type Events struct {
	subscribers []chan machine.Event

	sync.Mutex
}

// NewEvents initializes and returns the v1alpha1 runtime event stream.
func NewEvents(n int) *Events {
	e := &Events{
		subscribers: make([]chan machine.Event, 0, n),
	}

	return e
}

// Watch implements the Events interface.
func (e *Events) Watch(f runtime.WatchFunc) {
	ch := e.add()

	go func() {
		defer close(ch)
		defer e.delete(ch)

		f(ch)
	}()
}

// Publish implements the Events interface.
func (e *Events) Publish(msg proto.Message) {
	value, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("failed to marshal message: %v", err)

		return
	}

	event := machine.Event{
		Data: &any.Any{
			// In the future, we can publish `talos/runtime`, and
			// `talos/plugin/<plugin>` (or something along those lines) events.
			TypeUrl: fmt.Sprintf("talos/runtime/%s", proto.MessageName(msg)),
			Value:   value,
		},
	}

	e.Lock()
	defer e.Unlock()

	for _, sub := range e.subscribers {
		// drop the event if some subscriber is stuck
		// dropping is bad, but better than blocking event propagation
		select {
		case sub <- event:
		default:
		}
	}
}

func (e *Events) add() chan machine.Event {
	e.Lock()
	defer e.Unlock()

	ch := make(chan machine.Event, 100)

	e.subscribers = append(e.subscribers, ch)

	return ch
}

func (e *Events) delete(ch chan machine.Event) {
	e.Lock()
	defer e.Unlock()

	for i, sub := range e.subscribers {
		if sub == ch {
			l := len(e.subscribers)
			e.subscribers[i] = e.subscribers[l-1]
			e.subscribers[l-1] = nil
			e.subscribers = e.subscribers[:l-1]
		}
	}
}