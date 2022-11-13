package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"mall_srvs/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig

	EsClient *elastic.Client
)
