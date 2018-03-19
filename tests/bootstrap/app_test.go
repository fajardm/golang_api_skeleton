package bootstrap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"github.com/fajardm/example_blog_api/bootstrap"
	"net/http/httptest"
	"net/http"
)

func TestSetupRouter(t *testing.T) {
	// ensure SetupRouter is a  'function'
	assert.True(t, reflect.TypeOf(bootstrap.SetupRouter).Kind() == reflect.Func, "SetupRouter should a 'function'")

	// setup a router
	router := bootstrap.SetupRouter()

	// test endpoint ping
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "GET /ping should status OK")
	assert.Equal(t, "PONG", res.Body.String(), "GET /ping should return PONG")
}
