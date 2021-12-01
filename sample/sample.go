package sample

import "google.golang.org/grpc/resolver"

// 这里的sampleDriver，演示了一个的driver应当如何编写

type sampleDriver struct {
}

func (d *sampleDriver) GetName() string {
	return "dtm-sample-driver"
}

func (d *sampleDriver) RegisterService(url string, endpoint string) error {
	// 如果使用etcd，polaris等注册/发现组件的话，那么在这里，将endpoint注册到相应的url中
	// 这里的sample仅作为演示用，没有实际注册
	return nil
}

func (d *sampleDriver) RegisterGrpcResolver() {
	resolver.Register(&sampleBuilder{})
}
