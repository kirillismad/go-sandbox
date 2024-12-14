package service

import (
	"sandbox/utils"

	"google.golang.org/grpc/resolver"
)

var servers = []string{"localhost:50051", "localhost:50052"}

type staticResolverBuilder struct {
	addresses map[string][]string
}

func (b *staticResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	addresses := utils.Map(b.addresses[target.Endpoint()], func(addr string) resolver.Address {
		return resolver.Address{Addr: addr}
	})
	cc.UpdateState(resolver.State{Addresses: addresses})
	return &stubResolver{}, nil
}
func (*staticResolverBuilder) Scheme() string { return "static" }

type stubResolver struct{}

func (*stubResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (*stubResolver) Close()                                {}

func init() {
	resolver.Register(&staticResolverBuilder{
		addresses: map[string][]string{
			"service": servers,
		},
	})
}
