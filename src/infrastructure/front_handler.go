package infrastructure

import (
	"net/http"

	"recipe_api/src/interfaces/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type frontAPI struct {
	path    string
	method  string
	handler apiHandler
}

func newFrontAPI(path string, method string, handler apiHandler) *frontAPI {
	return &frontAPI{
		path:    path,
		method:  method,
		handler: handler,
	}
}

func handleFrontGroup(c *Controller, logger *zap.SugaredLogger, g *gin.RouterGroup) {
	logger = logger.With("router", "front")

	recipeController := controllers.NewHttpFrontRecipeController(c)

	api := g.Group("/v1")

	get := http.MethodGet
	post := http.MethodPost
	patch := http.MethodPatch
	delete := http.MethodDelete

	apis := []*frontAPI{
		newFrontAPI("/recipes", get, recipeController.FindAll),
		newFrontAPI("/recipes", post, recipeController.Create),
		newFrontAPI("/recipes/:recipe_id", get, recipeController.FindById),
		newFrontAPI("/recipes/:recipe_id", patch, recipeController.Update),
		newFrontAPI("/recipes/:recipe_id", delete, recipeController.Delete),
	}
	for _, a := range apis {

		handler := a.handler

		h := func(ctx *gin.Context) {
			status, res, err := handler(ctx)
			if err != nil {
				logger.Warn(err)
			}
			ctx.JSON(status, res)
		}
		api.Handle(a.method, a.path, h)
	}
}
