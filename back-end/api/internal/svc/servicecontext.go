package svc

import (
	"back-end/api/internal/config"
	"back-end/api/internal/middleware"
	"back-end/api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserModel           model.UserTblModel
	UserTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserModel:           model.NewUserTblModel(sqlx.NewSqlConn(c.Database.DataName, c.Database.DataSource)),
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
	}
}
