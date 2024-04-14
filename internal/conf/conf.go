package conf

import (
	"rose/common/net/chttp"
)

type Conf struct {
	Server *chttp.Config `yaml:"server"`
}
