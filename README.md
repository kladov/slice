Simple example:

```go
package main

import (
	"fmt"

	"github.com/kladov/slice"
)

var allowed = slice.NewStringSlice([]string{"foo", "bar"})

func main() {

	// ... do some business logic

	if allowed.Has("bar") {
		fmt.Println("bar is allowed")
	}

	// ... do some business logic

	if !allowed.Has("baz") {
		fmt.Printf("baz is not allowed, allowed values: %v", allowed.Values())
	}
}
```