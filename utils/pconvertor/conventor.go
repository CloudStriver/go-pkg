package pconvertor

import (
	"github.com/CloudStriver/go-pkg/utils/pagination"
	"github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
)

func StructToJsonString(a any) string {
	data, err := sonic.Marshal(a)
	if err != nil {
		logx.Errorf("Json Marshal异常[%v]\n", err)
		return ""
	}
	return string(data)
}

func PaginationOptionsToModelPaginationOptions(options *basic.PaginationOptions) *pagination.PaginationOptions {
	return &pagination.PaginationOptions{
		Limit:     options.Limit,
		Offset:    options.Offset,
		Backward:  options.Backward,
		LastToken: options.LastToken,
	}
}
