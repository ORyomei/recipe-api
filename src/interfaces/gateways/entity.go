package gateways

import (
	"time"
)

// Model is base model
type Model struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

// StringIDModel is base model with string id
type StringIDModel struct {
	ID        string    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}

// IDLessModel is base model without id
type IDLessModel struct {
	CreatedAt time.Time `gorm:"-"`
	UpdatedAt time.Time `gorm:"-"`
}
