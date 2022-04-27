package main

import (
	"mvc/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	app.StartRoute()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	//router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
