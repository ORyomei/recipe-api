package usecase

import (
	"context"
	"fmt"

	"recipe_api/src/domain/entity"
	"recipe_api/src/interfaces/gateways"
	"recipe_api/src/openapi/front/schema"
)

type FrontRecipeUsecase struct{}

func NewFrontRecipeUsecase() *FrontRecipeUsecase {
	return &FrontRecipeUsecase{}
}

func (*FrontRecipeUsecase) FindAll(ctx context.Context, mySql gateways.MySql) ([]entity.Recipe, error) {
	return gateways.FindAllRecipes(mySql.GetQuerier())
}

func (*FrontRecipeUsecase) Create(ctx context.Context, mySql gateways.MySql, body schema.PostRecipeRequestBody) (*entity.Recipe, error) {
	r, err := gateways.CreateRecipe(mySql.GetQuerier(), &entity.Recipe{
		Title:       body.Title,
		MakingTime:  body.MakingTime,
		Serves:      body.Serves,
		Ingredients: body.Ingredients,
		Cost:        body.Cost,
	})
	if err != nil {
		return nil, fmt.Errorf("gateways.Create failed: %w", err)
	}
	return r, nil
}

func (*FrontRecipeUsecase) FindById(ctx context.Context, mySql gateways.MySql, id uint64) (*entity.Recipe, error) {
	return gateways.FindRecipeById(mySql.GetQuerier(), id)
}

func (*FrontRecipeUsecase) Update(ctx context.Context, mySql gateways.MySql, body schema.PatchRecipeRequestBody, id uint64) (*entity.Recipe, error) {
	r, err := gateways.UpdateRecipe(mySql.GetQuerier(), &entity.Recipe{
		Title:       body.Title,
		MakingTime:  body.MakingTime,
		Serves:      body.Serves,
		Ingredients: body.Ingredients,
		Cost:        body.Cost,
	}, id)
	if err != nil {
		return nil, fmt.Errorf("gateways.Update failed: %w", err)
	}
	return r, nil
}

func (*FrontRecipeUsecase) Delete(ctx context.Context, mySql gateways.MySql, id uint64) error {
	err := gateways.DeleteRecipe(mySql.GetQuerier(), id)
	if err != nil {
		return fmt.Errorf("gateways.Delete failed: %w", err)
	}
	return nil
}
