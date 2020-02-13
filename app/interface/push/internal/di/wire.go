// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"chat/app/interface/push/internal/dao"
	"chat/app/interface/push/internal/server/http"
	"chat/app/interface/push/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, NewApp))
}
