package siotest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func NewGinReq(t *testing.T) *gin.Context {
	t.Helper()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}
