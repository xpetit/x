package x

import "net/http"

func init() {
	http.DefaultTransport.(*http.Transport).DisableCompression = true
}
