package http

import "net/http"

type Requestor interface {
	Do(req *http.Request) (*http.Response, error)
}

// Request ..
type Request struct{}

// NewRequestService func acts as a constructor.
func NewRequestService() Requestor {
	return &Request{}
}

// Do func is to make a REST call with the *http.Request passed to it.
func (r *Request) Do(req *http.Request) (*http.Response, error) {
	var client = newClient()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// newClient creates the new *http.Client
func newClient() *http.Client {
	return &http.Client{}
}
