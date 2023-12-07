package schema

import "recipe_api/src/domain/entity"

func NewIdLessRecipe(r entity.Recipe) *IdLessRecipe {
	return &IdLessRecipe{
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
	}
}

func NewPostedRecipe(r entity.Recipe) *PostedRecipe {
	return &PostedRecipe{
		Id:          r.Id,
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func NewRecipe(r entity.Recipe) *Recipe {
	return &Recipe{
		Id:          r.Id,
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
	}
}
