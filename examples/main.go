package main

import (
	"fmt"

	"github.com/coreservice-io/UHub"
)

const testKind UHub.Kind = 1

type testEvent string

func (e testEvent) Kind() UHub.Kind {
	return testKind
}

const testKind2 UHub.Kind = 2

type testEvent2 string

func (e testEvent2) Kind() UHub.Kind {
	return testKind2
}

func main() {

	var h UHub.Hub

	cancel := h.Subscribe(testKind, func(e UHub.Event) {
		fmt.Println("sub1")
		fmt.Println(string(e.(testEvent)))
	})
	cancel()

	h.Subscribe(testKind, func(e UHub.Event) {
		fmt.Println("sub2")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind, func(e UHub.Event) {
		fmt.Println("sub3")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind2, func(e UHub.Event) {
		fmt.Println("sub4")
		fmt.Println(string(e.(testEvent2)))
	})

	h.Publish(testEvent("this is example"))
	h.Publish(testEvent2("this is example2"))

}
