package upgrader

import (
	"net/http"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 10,
	}
)
