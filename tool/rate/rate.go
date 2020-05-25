package rate

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/hashicorp/vic/lib/apiservers/service/restapi/handlers/errors"
	"golang.org/x/time/rate"
)

//限流中间件
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.NewError(429, "too many request")
			}
			return next(ctx, request)
		}
	}
}
