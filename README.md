# Mini WebSite Render

MWR is Niubi!

![LICENSE](https://img.shields.io/badge/license-MWR-blue.svg)

## Quick Start 

```shell
go get github.com/mwr-wiki/mini-website-render
```

## Hello World

```go
package main

import (
	"fmt"
	"github.com/mwr-wiki/mini-website-render"
)

func main() {
	mwr := new(mini_website_render.Mwr)
	fmt.Print("Service Started")
	mwr.Run()
}
```