// +build appengine

package httpinit

import (
	"github.com/zenazn/goji"
	"net/http"
)

func Init() {
	http.Handle("/", goji.DefaultMux)
}
