package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/bitverb/x/errors"
)

func TestAny(t *testing.T) {
	type HelloRequest struct {
	}
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	OkJson[HelloRequest](c, HelloRequest{})
	Err[interface{}](c, errors.New(http.StatusOK, 1, ""), nil)
}
