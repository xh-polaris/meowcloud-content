//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowcloud-content/biz/adaptor"
)

func NewUserServerImpl() (*adaptor.ContentServerImpl, error) {
	wire.Build(
		wire.Struct(new(adaptor.ContentServerImpl), "*"),
		AllProvider,
	)
	return nil, nil
}
