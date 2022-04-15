package dtmdriver

// HTTPDriver is the interface to do http service register and discover
type HTTPDriver interface {
	// GetName return the name of the driver
	GetName() string
	// Init will init client
	Init(registryType string, host string, options string) error
	// ResolveURL register the http resolver to handle custom url
	ResolveURL(url string) (string, error)
	// RegisterService register dtm endpoint to service
	RegisterService(service string, endpoint string) error
}

// HTTPClient is the interface to do http service call.
type HTTPClient interface {
	// ResolveURL register the http resolver to handle custom url
	ResolveURL(url string) (string, error)
	// RegisterService register dtm endpoint to service
	RegisterService(service string, endpoint string) error
}
