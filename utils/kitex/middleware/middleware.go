package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"

	"github.com/CloudStriver/go-pkg/utils/util"
	"github.com/CloudStriver/go-pkg/utils/util/log"
)

var (
	LogMiddleware = func(name string) endpoint.Middleware {
		return func(handler endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, req, resp interface{}) error {
				err := handler(ctx, req, resp)
				log.CtxInfo(ctx, "[%s RPC Request] req=%s, resp=%s, err=%v", name, util.JSONF(req), util.JSONF(resp), err)
				return err
			}
		}
	}
)
