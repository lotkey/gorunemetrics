// Package http wraps the standard http library with interfaces that allow for
// injecting mocks for testing.
package http

import (
	"io"
	"net/http"
	"net/url"
)

type (
	Client interface {
		CloseIdleConnections()
		Do(req *http.Request) (*http.Response, error)
		Get(url string) (resp *http.Response, err error)
		Head(url string) (resp *http.Response, err error)
		Post(
			url, contentType string,
			body io.Reader,
		) (resp *http.Response, err error)
		PostForm(url string, data url.Values) (resp *http.Response, err error)
	}
)
