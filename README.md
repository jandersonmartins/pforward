# pforward

Simple port forward

## Install

```sh
$ go get -u github.com/jandersonmartins/pforward@v0.0.1
```

## Usage

```go
package main

import "github.com/jandersonmartins/pforward"

func main() {
	cnf := pforward.ForwardCnf{ServerPort: 8001,
		DestHost: "localhost",
		DestPort: 4000}
        
	pforward.Forward(cnf)
}
```