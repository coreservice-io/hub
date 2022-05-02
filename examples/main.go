package main

import (
	"fmt"

	"github.com/coreservice-io/hub"
)

const testKind hub.Kind = 1

type testEvent string

func (e testEvent) Kind() hub.Kind {
	return testKind
}

const testKind2 hub.Kind = 2

type testEvent2 string

func (e testEvent2) Kind() hub.Kind {
	return testKind2
}

func main() {

	var h hub.Hub

	cancel := h.Subscribe(testKind, func(e hub.Event) {
		fmt.Println("sub1")
		fmt.Println(string(e.(testEvent)))
	})
	cancel()

	h.Subscribe(testKind, func(e hub.Event) {
		fmt.Println("sub2")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind, func(e hub.Event) {
		fmt.Println("sub3")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind2, func(e hub.Event) {
		fmt.Println("sub4")
		fmt.Println(string(e.(testEvent2)))
	})

	h.Publish(testEvent("this is example"))
	h.Publish(testEvent2("this is example2"))

}
