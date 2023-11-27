package infrastructure

import (
	"context"
	"fmt"

	"recipe_api/src/interfaces/gateways"
	"recipe_api/src/util"

	"go.uber.org/zap"
)

type Controller struct {
	conf   *util.Config
	mySql  *MySql
	logger *zap.SugaredLogger
}

func NewController(config *util.Config, logger *zap.SugaredLogger) (*Controller, error) {
	mySql, err := OpenMySql(config.MySql)
	if err != nil {
		return nil, fmt.Errorf("OpenMySql failed: %w", err)
	}

	return &Controller{
		conf:   config,
		mySql:  mySql,
		logger: logger,
	}, nil
}

func (c *Controller) GetMySql(ctx context.Context) gateways.MySql {
	return c.mySql.WithContext(ctx)
}

func (c *Controller) GetLogger() *zap.SugaredLogger {
	return c.logger
}

func (c *Controller) GetConfig() *util.Config {
	return c.conf
}
