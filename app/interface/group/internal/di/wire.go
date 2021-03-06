// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"chat/app/interface/group/internal/dao"
	"chat/app/interface/group/internal/server"
	"chat/app/interface/group/internal/service"
	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, server.New, NewApp))
}
