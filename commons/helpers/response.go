package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// interface for gin.Context JSON
type IContext interface {
	JSON(code int, obj interface{})
}

type Response struct {
	status     string
	HttpStatus int
	Data       interface{}
	Message    string
	Code       string
}

func (res *Response) Success(c IContext) {
	res.status = "success"

	if res.HttpStatus == 0 {
		res.HttpStatus = http.StatusOK
	}

	c.JSON(res.HttpStatus, gin.H{
		"status": res.status,
		"data":   res.Data,
	})
}

func (res *Response) Fail(c IContext) {
	res.status = "fail"

	if res.HttpStatus == 0 {
		res.HttpStatus = http.StatusBadRequest
	}

	c.JSON(res.HttpStatus, gin.H{
		"status": res.status,
		"data":   res.Data,
	})
}

func (res *Response) Error(c IContext) {
	res.status = "error"

	if res.HttpStatus == 0 {
		res.HttpStatus = http.StatusInternalServerError
	}

	if res.Code != "" {
		c.JSON(res.HttpStatus, gin.H{
			"status":  res.status,
			"message": res.Message,
			"code":    res.Code,
		})
	} else {
		c.JSON(res.HttpStatus, gin.H{
			"status":  res.status,
			"message": res.Message,
		})
	}
}
