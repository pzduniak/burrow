burrow
======

Burrow is a Graval-based embedded FTP server library. How to use:

```go
package main

import (
	"github.com/pzduniak/burrow"
	"github.com/pzduniak/graval"
	"log"
)

func main() {
  // /home/test is the main server path
  // The second parameter is the authentication function
	factory := &burrow.PomFTPFactory{"/home/test", func(username string, password string) bool {
		if username == "test" && password == "1234" {
			return true
		}

		return false
	}}

  // Burrow only works with my fork of Graval, because of non-backwards-compatible patches
  // that allow support of large files
	ftpServer := graval.NewFTPServer(&graval.FTPServerOpts{Factory: factory, Port: 21})
	err := ftpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

```
