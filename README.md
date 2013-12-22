burrow
======

Burrow is a Graval-based embedded FTP server library. How to use:

```go
package main

import (
	"github.com/pzduniak/burrow"
	"log"
)

func main() {
	err := burrow.NewServer(burrow.Config{
		HomePath: "/home/test",
		Authenticate: func(username string, password string) bool {
			if username == "test" && password == "1234" {
				return true
			}

			return false
		},
		Port: 21,
	}).Listen()

	if err != nil {
		log.Printf("Error while listening: %s", err)
	}
}
```
