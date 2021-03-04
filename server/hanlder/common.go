package hanlder

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

func Success(r *ghttp.Request, message interface{}) {
	if value, ok := message.(string); ok {
		r.Response.WriteJsonExit(g.Map{
			"status":  200,
			"message": value,
		})
	}
	r.Response.WriteJsonExit(message)
}

func Failed(r *ghttp.Request, err interface{}, code ...interface{}) {
	var status int
	if len(code) > 0 {
		status = code[0].(int)
	} else {
		status = 500
	}

	var message = "failed"

	if v, ok := err.(*gvalid.Error); ok {
		message = v.FirstString()
		status = 422
	} else if v, ok := err.(error); ok {
		message = v.Error()
	} else if value, ok := err.(string); ok {
		message = value
	}

	r.Response.WriteHeader(status)
	r.Response.WriteJsonExit(g.Map{
		"status":  status,
		"message": message,
	})
}
