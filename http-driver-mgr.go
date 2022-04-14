package dtmdriver

import "fmt"

var (
	httpDrivers = map[string]HttpDriver{}
	httpCurrent HttpDriver
)

// RegisterHttp used by each driver writer to register the driver to system
func RegisterHttp(driver HttpDriver) {
	httpDrivers[driver.GetName()] = driver
}

// UseHttp called by users who want to use dtmgrpc.
func UseHttp(name string) error {
	v := httpDrivers[name]
	if v == nil {
		return fmt.Errorf("no dtm driver with name: %s has been registered", name)
	}
	if httpCurrent != nil && v != httpCurrent {
		return fmt.Errorf("use has been called previously with name: %s different from: %s", httpCurrent.GetName(), name)
	}
	if httpCurrent == nil {
		httpCurrent = v
	}
	return nil
}
