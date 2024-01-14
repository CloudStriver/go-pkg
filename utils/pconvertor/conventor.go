package pconvertor

import (
	"context"
	"github.com/CloudStriver/go-pkg/utils/pagination"
	"github.com/CloudStriver/go-pkg/utils/util/log"
	"github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	"github.com/bytedance/sonic"
)

func StructToJsonString(ctx context.Context, a any) string {
	data, err := sonic.Marshal(a)
	if err != nil {
		log.CtxError(ctx, "sonic Marshal异常[%v]\n", err)
		return ""
	}
	return string(data)
}

func JsonStringToStruct(ctx context.Context, a any, data []byte) {
	err := sonic.Unmarshal(data, a)
	if err != nil {
		log.CtxError(ctx, "sonic Unmarshal异常[%v]\n", err)
	}
}

func PaginationOptionsToModelPaginationOptions(options *basic.PaginationOptions) *pagination.PaginationOptions {
	if options == nil {
		return &pagination.PaginationOptions{}
	}
	return &pagination.PaginationOptions{
		Limit:     options.Limit,
		Offset:    options.Offset,
		Backward:  options.Backward,
		LastToken: options.LastToken,
	}
}
