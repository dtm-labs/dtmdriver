package dtmdriver

import "fmt"

var (
	grpcDrivers = map[string]Driver{}
	grpcCurrent Driver
)

// Register used by each driver writer to register the driver to system
func Register(driver Driver) {
	grpcDrivers[driver.GetName()] = driver
}

// Use called by users who want to use dtmgrpc.
func Use(name string) error {
	v := grpcDrivers[name]
	if v == nil {
		return fmt.Errorf("no dtm driver with name: %s has been registered", name)
	}
	if grpcCurrent != nil && v != grpcCurrent {
		return fmt.Errorf("use has been called previously with name: %s different from: %s", grpcCurrent.GetName(), name)
	}
	if grpcCurrent == nil {
		grpcCurrent = v
		v.RegisterGrpcResolver()
	}
	return nil
}

// GetDriver called by dtm
func GetDriver() Driver {
	if grpcCurrent == nil {
		return grpcDrivers["default"]
	}
	return grpcCurrent
}
