# hub
universal subscriber and publisher


```go

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

func main() {

	var h hub.Hub

	h.Subscribe(testKind, func(e hub.Event) {
		fmt.Println("subcallback")
		fmt.Println(string(e.(testEvent)))
	})
	h.Publish(testEvent("this is example"))

}


```
