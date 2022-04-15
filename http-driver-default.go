package dtmdriver

import "log"

type defaultHTTPDriver struct {
	client HTTPClient
}

func (d *defaultHTTPDriver) GetName() string {
	return "default"
}

func (d *defaultHTTPDriver) Init(registryType string, host string, options string) error {
	d.client = &defaultHTTPClient{}
	return nil
}

func (d *defaultHTTPDriver) ResolveURL(url string) (string, error) {
	log.Printf("d: %v", d)
	log.Printf("d.client: %v", d.client)
	log.Printf("d.client.ResolveURL: %v", d.client.ResolveURL)
	return d.client.ResolveURL(url)
}

func (d *defaultHTTPDriver) RegisterService(target string, endpoint string) error {
	return d.client.RegisterService(target, endpoint)
}

type defaultHTTPClient struct{}

func (d *defaultHTTPClient) ResolveURL(url string) (string, error) {
	return url, nil
}

func (d *defaultHTTPClient) RegisterService(target string, endpoint string) error {
	return nil
}

func init() {
	defaultDriver := &defaultHTTPDriver{}
	_ = defaultDriver.Init("", "", "") // default driver init need no params
	RegisterHttp(&defaultHTTPDriver{})
}
