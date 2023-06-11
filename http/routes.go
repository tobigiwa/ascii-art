package http

import "net/http"

func Routes() http.Handler {
	getRoute := http.HandlerFunc(GET)
	postRoute := http.HandlerFunc(POST)

	mux := http.NewServeMux()
	mux.Handle("/", HttpMethod(http.MethodGet, getRoute))
	mux.Handle("/ascii-art", HttpMethod(http.MethodPost, postRoute))

	return LogRequest(Recover(mux))
}
