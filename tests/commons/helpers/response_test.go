package helpers

import (
	"github.com/stretchr/testify/mock"
	"testing"
	"github.com/fajardm/example_blog_api/commons/helpers"
	"github.com/gin-gonic/gin"
)

// implement IContext for testing
type impContext struct {
	mock.Mock
}

func (c *impContext) JSON(code int, obj interface{}) {
	c.Called(code, obj)
}

func TestSuccess(t *testing.T) {
	impCtx := new(impContext)
	impCtx.On("JSON", 200, gin.H{
		"status": "success",
		"data":   "data",
	})

	response := helpers.Response{
		Data: "data",
	}
	response.Success(impCtx)
}

func TestFail(t *testing.T) {
	impCtx := new(impContext)
	impCtx.On("JSON", 400, gin.H{
		"status": "fail",
		"data":   "data",
	})

	response := helpers.Response{
		Data: "data",
	}
	response.Fail(impCtx)
}

func TestError(t *testing.T) {
	impCtx := new(impContext)
	impCtx.On("JSON", 500, gin.H{
		"status":  "error",
		"message": "message",
	})

	response := helpers.Response{
		Message: "message",
	}
	response.Error(impCtx)
}
