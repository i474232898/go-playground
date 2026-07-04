# Fan-In

Implement a `fanIn` function that merges multiple input channels into a single output channel.

## Task

Given two or more `<-chan int` channels (produced by a `generator` function), implement `fanIn` so that values from all input channels are forwarded to a single output channel concurrently. The output channel should be closed once all input channels are drained.

**Requirements:**

- Values from all input channels must be multiplexed onto one output channel
- Each input channel must be read concurrently (one goroutine per channel)
- The function signature should be: `func fanIn(chans ...<-chan int) <-chan int`

## 🔧 Example Stub

```go
package main

import (
	"fmt"
)
var p = fmt.Println

func fanIn(chans ...<-chan int) <-chan int {
	// your code here
}

func main() {
	// your code here
}
```
