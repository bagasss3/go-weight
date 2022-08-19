package model

import (
	"context"
	"time"
)

type Weight struct {
	ID         int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MaxWeight  int32     `json:"max_weight"`
	MinWeight  int32     `json:"min_weight"`
	Difference int32     `json:"difference"`
	CreatedAt  time.Time `json:"created_at"`
}

type WeightInput struct {
	MaxWeight int32  `json:"max_weight" validate:"required,min=1"`
	MinWeight int32  `json:"min_weight" validate:"required,min=1,ltfield=MaxWeight"`
	CreatedAt string `json:"created_at" validate:"required"`
}

type WeightRepository interface {
	Store(ctx context.Context, weight *Weight) error
	GetAllWeights(ctx context.Context) ([]*Weight, error)
	GetWeightByID(ctx context.Context, id int64) (*Weight, error)
	UpdateWeight(ctx context.Context, id int64, weight *Weight) error
	DeleteWeight(ctx context.Context, id int64) (bool, error)
}

type WeightController interface {
	Create(ctx context.Context, req *WeightInput) (*Weight, error)
	ReadWeights(ctx context.Context) ([]*Weight, error)
	ShowWeight(ctx context.Context, id int64) (*Weight, error)
	UpdateWeight(ctx context.Context, id int64, req *WeightInput) (*Weight, error)
	DeleteWeight(ctx context.Context, id int64) (bool, error)
}

func (w *WeightInput) Validate() error {
	return validate.Struct(w)
}
