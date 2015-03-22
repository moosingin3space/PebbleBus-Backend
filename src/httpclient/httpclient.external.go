// +build !appengine

package httpclient

import (
	"net/http"
)

func Client(r *http.Request) *http.Client {
	return &http.Client{}
}
