# UHub
universal subscriber and publisher


```go

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

func main() {

	var h UHub.Hub

	h.Subscribe(testKind, func(e UHub.Event) {
		fmt.Println("subcallback")
		fmt.Println(string(e.(testEvent)))
	})
	h.Publish(testEvent("this is example"))

}


```
