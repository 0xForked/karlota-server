package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessRespond message
type SuccessRespond struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"OK"`
	Data   interface{} `json:"data"`
}

// ErrorRespond message
type ErrorRespond struct {
	Code   int    `json:"code" example:"400"`
	Status string `json:"status" example:"Bad Request"`
	Data   string `json:"data"`
}

// ValidationErrorRespond message
type ValidationErrorRespond struct {
	Code   int         `json:"code" example:"422"`
	Status string      `json:"status" example:"Unprocessable Entity"`
	Data   interface{} `json:"data"`
}

// ServerErrorRespond message
type ServerErrorRespond struct {
	Code   int    `json:"code" example:"500"`
	Status string `json:"status" example:"Internal Server Error"`
	Data   string `json:"data"`
}

func NewHttpRespond(context *gin.Context, code int, data interface{}) {
	if code == http.StatusOK || code == http.StatusCreated {
		context.JSON(code, SuccessRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   data,
		})

		return
	}

	if code == http.StatusUnprocessableEntity {
		context.JSON(code, ValidationErrorRespond{
			Code:   code,
			Status: http.StatusText(code),
			Data:   data,
		})

		return
	}

	msg := func() string {
		switch {
		case data != nil:
			return data.(string)
		case code == http.StatusBadRequest:
			return "something went wrong with the request"
		default:
			return "something went wrong with the server"
		}
	}()

	context.JSON(code, ErrorRespond{
		Code:   code,
		Status: http.StatusText(code),
		Data:   msg,
	})

	return
}
