package iptools

import (
	"errors"
)

// Result represents the result of an IP Lookup
type Result struct {
	Error error

	V4 *string
	V6 *string
}

// Lookup the current machine's public IPv4 and IPv6 address
func Lookup() <-chan *Result {
	r := make(chan *Result)

	go func() {
		defer close(r)
		res := &Result{}

		v4task, v6task := getAddress("v4"), getAddress("v6")
		v4ptr, v6ptr := <-v4task, <-v6task

		if v4ptr.Data == nil && v6ptr.Data == nil {
			res.Error = errors.New("both IPv4 and IPv6 cannot be found")
			r <- res

			return
		}

		res.V4 = v4ptr.Data
		res.V6 = v6ptr.Data

		r <- res
	}()

	return r
}
