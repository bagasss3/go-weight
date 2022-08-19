package controller

import (
	"context"
	"errors"
	"time"

	"github.com/bagasss3/go-weight/internal/config"
	"github.com/bagasss3/go-weight/internal/model"
	"github.com/sirupsen/logrus"
)

type weightController struct {
	weightRepo model.WeightRepository
}

func NewWeightController(weightRepo model.WeightRepository) model.WeightController {
	return &weightController{
		weightRepo: weightRepo,
	}
}

func (w *weightController) Create(ctx context.Context, req *model.WeightInput) (*model.Weight, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"req": req,
	})

	if err := req.Validate(); err != nil {
		logger.Error(err)
		return nil, err
	}

	parse, err := time.Parse(config.LayoutISO, req.CreatedAt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	weight := &model.Weight{
		MaxWeight:  req.MaxWeight,
		MinWeight:  req.MinWeight,
		Difference: req.MaxWeight - req.MinWeight,
		CreatedAt:  parse,
	}
	if err := w.weightRepo.Store(ctx, weight); err != nil {
		logger.Error(err)
		return nil, err
	}
	return weight, nil
}

func (w *weightController) ReadWeights(ctx context.Context) ([]*model.Weight, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})

	weights, err := w.weightRepo.GetAllWeights(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return weights, nil
}

func (w *weightController) ShowWeight(ctx context.Context, id int64) (*model.Weight, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	weight, err := w.weightRepo.GetWeightByID(ctx, id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if weight == nil {
		return nil, errors.New("Weight not found")
	}
	return weight, nil
}

func (w *weightController) UpdateWeight(ctx context.Context, id int64, req *model.WeightInput) (*model.Weight, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	parse, err := time.Parse(config.LayoutISO, req.CreatedAt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	weight, err := w.ShowWeight(ctx, id)
	if err != nil {
		return nil, err
	}

	weight.MaxWeight = req.MaxWeight
	weight.MinWeight = req.MinWeight
	weight.Difference = req.MaxWeight - req.MinWeight
	weight.CreatedAt = parse

	err = w.weightRepo.UpdateWeight(ctx, id, weight)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return weight, nil

}

func (w *weightController) DeleteWeight(ctx context.Context, id int64) (bool, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	weight, err := w.ShowWeight(ctx, id)
	if err != nil {
		logger.Error(err)
		return false, err
	}

	isDeleted, err := w.weightRepo.DeleteWeight(ctx, weight.ID)
	if err != nil {
		logger.Error(err)
		return false, err
	}
	return isDeleted, nil
}
