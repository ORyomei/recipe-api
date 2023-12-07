package gateways

import (
	"recipe_api/src/domain/entity"

	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Recipe struct {
	Model
	Title       string `gorm:"size:256"`
	MakingTime  string `gorm:"size:256"`
	Serves      string `gorm:"size:256"`
	Ingredients string `gorm:"size:256"`
	Cost        int
}

func (r *Recipe) toEntity() *entity.Recipe {
	return &entity.Recipe{
		Id:          r.ID,
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func newRecipe(r *entity.Recipe) *Recipe {
	return &Recipe{
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
	}
}

func CreateRecipe(db *gorm.DB, Recipe *entity.Recipe) (*entity.Recipe, error) {
	r := newRecipe(Recipe)
	if err := db.Save(&r).Error; err != nil {
		return nil, err
	}
	return r.toEntity(), nil
}

func FindRecipeById(db *gorm.DB, id uint64) (*entity.Recipe, error) {
	var r Recipe
	err := db.Preload(clause.Associations).First(&r, id).Error
	if err != nil {
		return nil, err
	}
	return r.toEntity(), nil
}

func FindAllRecipes(db *gorm.DB) ([]entity.Recipe, error) {
	rs := []Recipe{}
	err := db.Find(&rs).Error

	if err != nil {
		return nil, err
	}
	recipes := []entity.Recipe{}
	for _, i := range rs {
		recipe := i.toEntity()
		recipes = append(recipes, *recipe)
	}
	return recipes, nil
}

func UpdateRecipe(db *gorm.DB, recipe *entity.Recipe, id uint64) (*entity.Recipe, error) {
	r := newRecipe(recipe)
	r.ID = id
	var r_ Recipe
	if err := db.Model(&r_).Clauses(clause.Returning{}).Save(r).Error; err != nil {
		return nil, fmt.Errorf("failed to save Item: %w", err)
	}
	return r_.toEntity(), nil
}

func DeleteRecipe(db *gorm.DB, id uint64) error {
	if err := db.Delete(&Recipe{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete Item: %w", err)
	}
	return nil
}
