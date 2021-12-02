package dtmdriver

import (
	"fmt"
	"strings"
)

type defaultDriver struct {
}

func (d *defaultDriver) GetName() string {
	return "default"
}

func (d *defaultDriver) RegisterGrpcResolver() {
}

func (d *defaultDriver) RegisterGrpcService(target string, endpoint string) error {
	return nil
}

func (d *defaultDriver) ParseServerMethod(url string) (svr string, method string, err error) {
	sep := strings.IndexByte(url, '/')
	if sep == -1 {
		return "", "", fmt.Errorf("bad url: '%s'. no '/' found", url)
	}
	return url[:sep], url[sep:], nil
}

func init() {
	Register(&defaultDriver{})
}
