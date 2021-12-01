package sample

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

type nopResolver struct {
	cc resolver.ClientConn
}

func (r *nopResolver) Close() {
}

func (r *nopResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

type sampleBuilder struct{}

func (s *sampleBuilder) Scheme() string { return "sample" }

func (s *sampleBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	endpoint := strings.ReplaceAll(target.Endpoint, s.Scheme()+"://", "")
	addrs := []resolver.Address{{Addr: endpoint}}

	if err := cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return nil, err
	}
	return &nopResolver{cc: cc}, nil
}
