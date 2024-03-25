package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	expect := "Hello, World!"
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	RootHandler(w, r)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	dataStr := string(data)
	assert.Nil(t, err)
	assert.Equal(t, expect, dataStr)
}
