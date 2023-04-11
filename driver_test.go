package driver

import (
	"testing"
)

func TestTaoDriver_RegisterGrpcService(t *testing.T) {

	// nacos
	target := "etcd://localhost:2379/dtmservice"
	endpoint := "localhost:8888"
	driver := new(taoDriver)
	if err := driver.RegisterService(target, endpoint); err != nil {
		t.Errorf("register etcd fail err :%+v", err)
	}

}
