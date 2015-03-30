// +build !appengine

package main

import (
	"github.com/zenazn/goji"
)

func main() {
	goji.Serve()
}
