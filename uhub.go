package uhub

import (
	"sync"
)

type Kind int64

type Event interface {
	Kind() Kind
}

type handler struct {
	f func(Event)
}

type Hub struct {
	subscribers map[Kind]map[string]handler
	rwmutex     sync.RWMutex
	counter     uint64
}

// Subscribe registers f for the event of a specific kind.
func (h *Hub) Subscribe(kind Kind, f func(Event)) (cancel func()) {
	cancelled := false
	h.rwmutex.Lock()
	if h.subscribers == nil {
		h.subscribers = make(map[Kind]map[string]handler)
	}
	if _, ok := h.subscribers[kind]; !ok {
		h.subscribers[kind] = make(map[string]handler)
	}
	rKey := randStr(32)
	for {
		if _, ok := h.subscribers[kind][rKey]; !ok {
			h.subscribers[kind][rKey] = handler{f: f}
			break
		}
		rKey = randStr(32)
	}
	h.counter++
	h.rwmutex.Unlock()

	cancel = func() {
		if !cancelled {
			h.rwmutex.Lock()
			cancelled = true
			delete(h.subscribers[kind], rKey)
			if len(h.subscribers[kind]) == 0 {
				delete(h.subscribers, kind)
			}
			h.counter--
			h.rwmutex.Unlock()
		}
	}
	return
}

// Publish an event to the subscribers.
func (h *Hub) Publish(e Event) {
	h.rwmutex.RLock()
	if handlers, ok := h.subscribers[e.Kind()]; ok {
		for _, h := range handlers {
			h.f(e)
		}
	}
	h.rwmutex.RUnlock()
}
