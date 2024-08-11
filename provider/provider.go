package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowcloud-content/biz/application/service"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
	album "github.com/xh-polaris/meowcloud-content/biz/infrastructure/mapper/album"
	photo "github.com/xh-polaris/meowcloud-content/biz/infrastructure/mapper/photo"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/stores/redis"
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	InfrastructureSet,
)

var ApplicationSet = wire.NewSet(
	service.AlbumSet,
	service.PhotoSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	redis.NewRedis,
	MapperSet,
)

var MapperSet = wire.NewSet(
	album.NewMongoMapper,
	photo.NewMongoMapper,
)
