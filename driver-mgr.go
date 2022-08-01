package dtmdriver

import "fmt"

var (
	drivers = map[string]Driver{}
	current Driver
)

// Register used by each driver writer to register the driver to system
func Register(driver Driver) {
	drivers[driver.GetName()] = driver
}

// Use called by users who want to use dtmgrpc.
func Use(name string) error {
	v := drivers[name]
	if v == nil {
		return fmt.Errorf("no dtm driver with name: %s has been registered", name)
	}
	if current != nil && v != current {
		return fmt.Errorf("use has been called previously with name: %s different from: %s", current.GetName(), name)
	}
	if current == nil {
		current = v
		v.RegisterAddrResolver()
	}
	return nil
}

// GetDriver called by dtm
func GetDriver() Driver {
	if current == nil {
		return drivers["default"]
	}
	return current
}
