package resolver

import (
	"google.golang.org/grpc/resolver"
)

// 定时，定期产生 address
type hackResolver struct {
	cc resolver.ClientConn
}

func init() {
	resolver.Register(NewBuilder())
}

// 主动更新
func (h hackResolver) ResolveNow(options resolver.ResolveNowOptions) {
	state := resolver.State{
		Addresses: []resolver.Address{
			resolver.Address{
				Addr:       "127.0.0.1:8082",
				ServerName: "test1",
				Attributes: nil,
			},
			resolver.Address{
				Addr:       "127.0.0.1:8081",
				ServerName: "test1",
				Attributes: nil,
			},
		},

		ServiceConfig: nil,
	}

	h.cc.UpdateState(state)
}

// 清理动作
func (h hackResolver) Close() {}

type hackResolveBuilder struct{}

func (h hackResolveBuilder) Build(
	target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return hackResolver{
		cc: cc,
	}, nil
}

func (h hackResolveBuilder) Scheme() string {
	return "hack"
}

func NewBuilder() hackResolveBuilder {
	return hackResolveBuilder{}
}
