# UHub
subscriber and publisher


```go

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

func main() {

	var h uhub.Hub

	h.Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("subcallback")
		fmt.Println(string(e.(testEvent)))
	})
	h.Publish(testEvent("this is example"))

}


```
