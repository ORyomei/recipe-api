package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
	"recipe_api/src/util"
)

func RunHttpServer() {
	logger, err := NewLogger()
	if err != nil {
		log.Printf("NewLogger failed: %s", err)
	}
	// logger = logger.With(
	// 	zap.String("app", "udp-server"),
	// )

	config := util.NewConfig()

	c, err := NewController(config, logger)

	engine := gin.New()

	engine.ContextWithFallback = true

	handleFrontGroup(c, logger, engine.Group("/front"))

	logger.Infof("running server: %s", "8080")
	if err := engine.Run(":8080"); err != nil {
		logger.Infof("finished running server: %s", err)
	}

}

type apiHandler func(*gin.Context) (int, interface{}, error)
