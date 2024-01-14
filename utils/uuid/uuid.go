package uuid

import (
	"context"
	"github.com/CloudStriver/go-pkg/utils/util/log"
	"github.com/gofrs/uuid"
	"github.com/zeromicro/go-zero/core/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"time"
)

func NewUuid(ctx context.Context) string {
	ctx, span := trace.TracerFromContext(ctx).Start(ctx, "uuid/new", oteltrace.WithTimestamp(time.Now()), oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer func() {
		span.End(oteltrace.WithTimestamp(time.Now()))
	}()
	uid, err := uuid.NewV4()
	if err != nil {
		log.CtxError(ctx, "uuid生成异常[%v]\n", err)
		return ""
	}
	return uid.String()
}
