# random

A package for random string. The intention is to provide a simple random string for transaction id for things like logging.

## Install

```
go get github.com/malcolm-davis/go-random
```

## Example

```go
package main

import (
	"fmt"

	"github.com/malcolm-davis/go-random"
)

func main() {
	// Print a random string of length 10
	fmt.Println(random.RandString(10))
}
```