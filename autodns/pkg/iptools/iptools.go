package iptools

// Result represents the result of an IP Lookup
type Result struct {
	Error error

	v4 *string
	v6 *string
}

// Lookup the current machine's public IPv4 and IPv6 address
func Lookup() <-chan *Result {
	r := make(chan *Result)

	go func() {
		defer close(r)
		res := &Result{}

		v4task, v6task := getAddress("v4"), getAddress("v6")
		v4ptr, v6ptr := <-v4task, <-v6task

		if v4ptr.Error != nil {
			res.Error = v4ptr.Error
			r <- res

			return
		}

		if v6ptr.Error != nil {
			res.Error = v6ptr.Error
			r <- res

			return
		}

		res.v4 = v4ptr.Data
		res.v6 = v6ptr.Data

		r <- res
	}()

	return r
}
