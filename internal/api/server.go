package api

import "net/http"

type API struct {
	s *http.Server
}

func NewAPI(addr string) API {
	return API{
		s: &http.Server{
			Handler: newHandler(),
			Addr:    addr,
		},
	}
}

func newHandler() http.Handler {
	mux := &http.ServeMux{}
	mux.Handle("POST /playlist", createPlaylistHandler())
	return mux
}

func (a *API) ListenAndServe() error {
	return a.s.ListenAndServe()
}
