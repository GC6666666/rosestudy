package conf

import (
	"rose/common/database"
	"rose/common/net/chttp"
)

type Conf struct {
	Server *chttp.Config    `yaml:"server"`
	DB     *database.Config `yaml:"db"`
}
