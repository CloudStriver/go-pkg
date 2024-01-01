package uuid

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"time"
)

func Newuuid(ctx context.Context) string {
	ctx, span := trace.TracerFromContext(ctx).Start(ctx, "uuid/new", oteltrace.WithTimestamp(time.Now()), oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer func() {
		span.End(oteltrace.WithTimestamp(time.Now()))
	}()
	uid, err := uuid.NewV4()
	if err != nil {
		logx.Errorf("uuid NewV4异常[%v]\n", err)
	}
	return uid.String()
}
