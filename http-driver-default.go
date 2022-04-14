package dtmdriver

type defaultHttpDriver struct {
}

func (d *defaultHttpDriver) GetName() string {
	return "default"
}

func (d *defaultHttpDriver) ResolveHttpURL(url string) (string, error) {
	return url, nil
}

func (d *defaultHttpDriver) RegisterHttpService(target string, endpoint string) error {
	return nil
}

func init() {
	RegisterHttp(&defaultHttpDriver{})
}
