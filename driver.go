package dtmdriver

import (
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc"
)

// Driver is the interface to do grpc service register and discover
type Driver interface {
	// GetName return the name of the driver
	GetName() string

	// RegisterAddrResolver will be called when driver used
	// for gRPC: register grpc resolver
	// for HTTP: add your http middleware
	RegisterAddrResolver()

	// RegisterService register dtm endpoint to target.
	// for both http and grpc
	RegisterService(target string, endpoint string) error

	// ParseServerMethod parse the gRPC url to server and method.
	// server will be passed to grpc.Dial, and method to grpc.ClientConn.invoke
	ParseServerMethod(uri string) (server string, method string, err error)
}

type middlewares struct {
	Grpc []grpc.UnaryClientInterceptor
	HTTP []func(c *resty.Client, r *resty.Request) error
}

// Middlewares is the middlewares for drivers to implement their function
var Middlewares = middlewares{
	Grpc: []grpc.UnaryClientInterceptor{},
	HTTP: []func(c *resty.Client, r *resty.Request) error{},
}
