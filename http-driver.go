package dtmdriver

// HttpDriver is the interface to do http service register and discover
type HttpDriver interface {
	// GetName return the name of the driver
	GetName() string
	// ResolveHttpURL register the http resolver to handle custom url
	ResolveHttpURL(url string) (string, error)
	// RegisterHttpService register dtm endpoint to target
	RegisterHttpService(target string, endpoint string) error
}
