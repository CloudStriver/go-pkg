package client

import (
	"context"
	"github.com/CloudStriver/go-pkg/utils/kitex/middleware"
	"github.com/CloudStriver/go-pkg/utils/util/log"
	"net"
	"strings"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

const (
	EnvHeader     = "X_XH_ENV"
	LaneHeader    = "X_Xh_LANE"
	magicEndpoint = "magic-host:magic-port"
)

var tracer = prometheus.NewClientTracer(":9091", "/client/metrics")

func NewClient[C any](fromName, toName string, fn func(fromName string, opts ...client.Option) (C, error)) C {
	cli, err := fn(
		fromName,
		client.WithHostPorts(func() []string {
			return []string{magicEndpoint}
		}()...),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithTracer(tracer),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fromName}),
		client.WithInstanceMW(middleware.LogMiddleware(toName)),
		client.WithLoadBalancer(&LoadBalancer{ServiceName: strings.ReplaceAll(toName, ".", "-")}),
	)
	if err != nil {
		log.Error("[NewClient], err=%v", err)
	}
	return cli
}

type LoadBalancer struct {
	ServiceName string
}

func (b *LoadBalancer) GetPicker(result discovery.Result) loadbalance.Picker {
	return &Picker{
		ServiceName: b.ServiceName,
		Instances:   result.Instances,
	}
}

func (b *LoadBalancer) Name() string {
	return "magic-name"
}

type Picker struct {
	ServiceName string
	Instances   []discovery.Instance
}

func (p *Picker) Next(ctx context.Context, _ interface{}) discovery.Instance {
	if len(p.Instances) != 0 && p.Instances[0].Address().String() != magicEndpoint {
		return p.Instances[0]
	}

	var host = p.ServiceName + ".cloudmind"

	// 选择基准环境
	env, ok := metainfo.GetPersistentValue(ctx, EnvHeader)
	if !ok {
		var md metadata.MD
		md, ok = metadata.FromIncomingContext(ctx)
		if ok && len(md[EnvHeader]) > 0 {
			env = md[EnvHeader][0]
		}
	}
	if ok && env == "test" {
		host += "-test"
	}

	// 检查泳道是否部署该服务
	lane, ok := metainfo.GetPersistentValue(ctx, LaneHeader)
	if !ok {
		var md metadata.MD
		md, ok = metadata.FromIncomingContext(ctx)
		if ok && len(md[LaneHeader]) > 0 {
			lane = md[LaneHeader][0]
		}
	}
	if ok && lane != "" {
		addr, err := net.ResolveTCPAddr("tcp", host+"-"+lane+".svc.cluster.local:8080")
		if err == nil {
			return &Instance{addr: addr}
		}
	}

	addr, err := net.ResolveTCPAddr("tcp", host+".svc.cluster.local:8080")
	if err == nil {
		return &Instance{addr: addr}
	}
	return nil
}

type Instance struct {
	addr net.Addr
}

func (i *Instance) Address() net.Addr {
	return i.addr
}

func (i *Instance) Weight() int {
	return 0
}

func (i *Instance) Tag(_ string) (value string, exist bool) {
	return "", false
}
