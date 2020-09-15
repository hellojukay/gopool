# gopool
an extremely simple goroutine pool

# example
```golang
package main

import (
	"time"

	"github.com/hellojukay/gopool/pool"
)

func main() {
	var p = pool.New(20)
	for {
		p.Run(func() {
			time.Sleep(time.Second * time.Duration(1))
		})
	}
}
```
