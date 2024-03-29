package middleware

import (
	"context"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/CloudStriver/go-pkg/utils/kitex/client"
)

var (
	EnvironmentMiddleware = func(ctx context.Context, c *app.RequestContext) {
		if env := c.Request.Header.Get("X-Xh-Env"); env != "" {
			fmt.Println(env)
			ctx = metainfo.WithPersistentValue(ctx, client.EnvHeader, env)
		}
		if lane := c.Request.Header.Get("X-Xh-Lane"); lane != "" {
			ctx = metainfo.WithPersistentValue(ctx, client.LaneHeader, lane)
		}
		c.Next(ctx)
	}
)
