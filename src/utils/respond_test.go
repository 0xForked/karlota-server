package utils_test

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNewHttpRespond(t *testing.T) {
	type args struct {
		context *gin.Context
		code    int
		data    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			//name: "success",
			//args: args{
			//	context: &gin.Context{},
			//	code:    200,
			//	data:    "success",
			//},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//utils.NewHttpRespond(tt.args.context, tt.args.code, tt.args.code)
		})
	}
}
