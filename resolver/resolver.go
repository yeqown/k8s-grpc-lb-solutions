package resolver

import (
	grpcresolver "google.golang.org/grpc/resolver"
)

// 定时，定期产生 address
type hackResolver struct {
	cc grpcresolver.ClientConn
}

func init() {
	grpcresolver.Register(NewBuilder())
}

// 主动更新
func (h hackResolver) ResolveNow(options grpcresolver.ResolveNowOptions) {
	state := grpcresolver.State{
		Addresses: []grpcresolver.Address{
			grpcresolver.Address{
				Addr:       "127.0.0.1:8082",
				ServerName: "test1",
				Attributes: nil,
			},
			grpcresolver.Address{
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
	target grpcresolver.Target, cc grpcresolver.ClientConn, opts grpcresolver.BuildOptions) (grpcresolver.Resolver, error) {
	r :=  hackResolver{
		cc: cc,
	}
	r.ResolveNow(grpcresolver.ResolveNowOptions{})
	return r, nil
}

func (h hackResolveBuilder) Scheme() string {
	return "hack"
}

func NewBuilder() hackResolveBuilder {
	return hackResolveBuilder{}
}
