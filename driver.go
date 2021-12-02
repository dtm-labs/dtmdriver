package dtmdriver

import "fmt"

// Driver interface to do service register and discover
type Driver interface {
	// GetName return the name of the driver
	GetName() string
	// RegisterGrpcResolver register the grpc resolver to handle custom scheme
	RegisterGrpcResolver()
	// RegisterGrpcService register dtm endpoint to target
	RegisterGrpcService(target string, endpoint string) error
	// ParseServerMethod parse the uri to server and method.
	// server will be passed to grpc.Dial, and method to grpc.ClientConn.invoke
	ParseServerMethod(uri string) (server string, method string, err error)
}

var (
	drivers = map[string]Driver{}
)

// Register used by each driver writer
func Register(driver Driver) {
	drivers[driver.GetName()] = driver
}

func MustGetDriver(name string) Driver {
	v := drivers[name]
	if v == nil {
		panic(fmt.Errorf("no dtm driver with name: %s has been registered", name))
	}
	return v
}
