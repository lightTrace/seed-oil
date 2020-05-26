package log

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"gokit/userserver/service"
)

//userService的日志中间件
func UserServiceLogger(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(service.UserRequest)
			logger.Log("method", r.Method, "event", "getUser")
			return next(ctx, request)
		}
	}
}
