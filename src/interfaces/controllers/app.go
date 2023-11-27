package controllers

import (
	"context"

	"recipe_api/src/interfaces/gateways"
	"recipe_api/src/util"

	"go.uber.org/zap"
)

type App interface {
	GetMySql(context.Context) gateways.MySql
	GetLogger() *zap.SugaredLogger
	GetConfig() *util.Config
}
