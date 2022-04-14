package dtmdriver

// Driver is the interface to do grpc service register and discover
type Driver interface {
	// GetName return the name of the driver
	GetName() string
	// RegisterGrpcResolver register the grpc resolver to handle custom scheme
	RegisterGrpcResolver()
	// RegisterGrpcService register dtm endpoint to target
	RegisterGrpcService(target string, endpoint string) error
	// ParseServerMethod parse the url to server and method.
	// server will be passed to grpc.Dial, and method to grpc.ClientConn.invoke
	ParseServerMethod(uri string) (server string, method string, err error)
}
