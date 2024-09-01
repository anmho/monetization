package test

import (
	"github.com/anmho/buy-me-a-boba/api"
	"net/http/httptest"
	"testing"
)

func MakeTestServer(t *testing.T) *httptest.Server {
	s := api.MakeServer()
	srv := httptest.NewServer(s)
	t.Cleanup(func() {
		srv.Close()
	})

	return srv
}
