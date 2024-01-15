package pconvertor

import (
	"github.com/CloudStriver/go-pkg/utils/pagination"
	"github.com/CloudStriver/go-pkg/utils/util/log"
	"github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	"github.com/bytedance/sonic"
	"reflect"
	"unsafe"
)

func StructToJsonString(a any) string {
	data, err := sonic.Marshal(a)
	if err != nil {
		log.Error("sonic Marshal异常[%v]\n", err)
		return ""
	}
	return Bytes2String(data)
}

func JsonStringToStruct(a any, data []byte) {
	err := sonic.Unmarshal(data, a)
	if err != nil {
		log.Error("sonic Unmarshal异常[%v]\n", err)
	}
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
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
