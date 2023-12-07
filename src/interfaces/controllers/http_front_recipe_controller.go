package controllers

import (
	"net/http"
	"recipe_api/src/openapi/front/schema"
	"recipe_api/src/usecase"

	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpFrontRecipeController struct {
	App
	Usecase usecase.FrontRecipeUsecase
}

func NewHttpFrontRecipeController(app App) *HttpFrontRecipeController {
	return &HttpFrontRecipeController{
		App: app,
	}
}

func (c *HttpFrontRecipeController) Create(ctx *gin.Context) (int, interface{}, error) {
	var body schema.PostRecipeRequestBody
	if e, err := BindJSON(ctx, &body); err != nil {
		return e.Status, e, err
	}

	r, err := c.Usecase.Create(ctx, c.GetMySql(ctx), body)
	if err != nil {
		e := NewInternalServerError()
		return e.Status, e, err
	}
	return http.StatusOK, schema.PostRecipeResponseBody{
		Recipe: &[]schema.PostedRecipe{*schema.NewPostedRecipe(*r)},
	}, nil
}

func (c *HttpFrontRecipeController) FindAll(ctx *gin.Context) (int, interface{}, error) {
	rs, err := c.Usecase.FindAll(ctx, c.GetMySql(ctx))
	if err != nil {
		e := NewInternalServerError()
		return e.Status, e, err
	}
	recipes := []schema.Recipe{}
	for _, r := range rs {
		recipes = append(recipes, *schema.NewRecipe(r))
	}
	return http.StatusOK, schema.GetRecipesResponseBody{
		Recipes: &recipes,
	}, nil
}

func (c *HttpFrontRecipeController) FindById(ctx *gin.Context) (int, interface{}, error) {
	recipeId, e, err := BindUintPathParam(ctx, "recipe_id")
	if err != nil {
		return e.Status, e, fmt.Errorf("BindUintPathParam failed: %w", err)
	}
	r, err := c.Usecase.FindById(ctx, c.GetMySql(ctx), uint64(recipeId))
	if err != nil {
		e := NewInternalServerError()
		return e.Status, e, err
	}
	m := "Recipe details by id"
	return http.StatusOK, schema.GetRecipeResponseBody{
		Message: &m,
		Recipe:  &[]schema.Recipe{*schema.NewRecipe(*r)},
	}, nil
}

func (c *HttpFrontRecipeController) Update(ctx *gin.Context) (int, interface{}, error) {
	recipeId, e, err := BindUintPathParam(ctx, "recipe_id")
	if err != nil {
		return e.Status, e, fmt.Errorf("BindUintPathParam failed: %w", err)
	}
	var body schema.PatchRecipeRequestBody
	if e, err := BindJSON(ctx, &body); err != nil {
		return e.Status, e, err
	}

	r, err := c.Usecase.Update(ctx, c.GetMySql(ctx), body, uint64(recipeId))
	if err != nil {
		e := NewInternalServerError()
		return e.Status, e, err
	}
	m := "Recipe successfully updated!"
	return http.StatusOK, schema.PatchRecipeResponseBody{
		Message: &m,
		Recipe:  &[]schema.IdLessRecipe{*schema.NewIdLessRecipe(*r)},
	}, nil

}

func (c *HttpFrontRecipeController) Delete(ctx *gin.Context) (int, interface{}, error) {
	recipeId, e, err := BindUintPathParam(ctx, "recipe_id")
	if err != nil {
		return e.Status, e, fmt.Errorf("BindUintPathParam failed: %w", err)
	}
	err = c.Usecase.Delete(ctx, c.GetMySql(ctx), uint64(recipeId))
	if err != nil {
		m := "No Recipe found"
		return http.StatusOK, schema.DeleteRecipeResponseBody{
			Message: &m,
		}, nil
	}
	m := "Recipe successfully removed!"
	return http.StatusOK, schema.DeleteRecipeResponseBody{
		Message: &m,
	}, nil
}
