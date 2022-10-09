package config

import "fmt"

var (
	v1 byte
	v2 byte
	v3 byte

	AppName  = "security-network"
	Protocol []byte
)

func init() {
	Protocol = make([]byte, 0)
	Protocol = append(Protocol, AppName...)
	Protocol = append(Protocol, v1, v2, v3)
}

func CheckVersion(vv1, vv2, vv3 byte) error {
	if vv1 != v1 || vv2 != v2 || vv3 != v3 {
		return fmt.Errorf("版本错误, 请使用%d.%d.%d", v1, v2, v3)
	}
	return nil
}
