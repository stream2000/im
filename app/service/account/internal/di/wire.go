// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"chat/app/service/account/internal/dao"
	"chat/app/service/account/internal/server/grpc"
	"chat/app/service/account/internal/server/http"
	"chat/app/service/account/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
