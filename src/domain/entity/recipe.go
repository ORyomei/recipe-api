package entity

import "time"

type Recipe struct {
	Id          uint64
	Title       string
	MakingTime  string
	Serves      string
	Ingredients string
	Cost        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
