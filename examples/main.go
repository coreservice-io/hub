package main

import (
	"fmt"

	uhub "github.com/coreservice-io/UHub"
)

const testKind uhub.Kind = 1

type testEvent string

func (e testEvent) Kind() uhub.Kind {
	return testKind
}

const testKind2 uhub.Kind = 2

type testEvent2 string

func (e testEvent2) Kind() uhub.Kind {
	return testKind2
}

func main() {

	var h uhub.Hub

	cancel := h.Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("sub1")
		fmt.Println(string(e.(testEvent)))
	})
	cancel()

	h.Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("sub2")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("sub3")
		fmt.Println(string(e.(testEvent)))
	})

	h.Subscribe(testKind2, func(e uhub.Event) {
		fmt.Println("sub4")
		fmt.Println(string(e.(testEvent2)))
	})

	h.Publish(testEvent("this is example"))
	h.Publish(testEvent2("this is example2"))

}
