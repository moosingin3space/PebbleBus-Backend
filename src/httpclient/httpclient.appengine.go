// +build appengine

package httpclient

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
)

func Client(r *http.Request) *http.Client {
	c := appengine.NewContext(r)
	return urlfetch.Client(c)
}
