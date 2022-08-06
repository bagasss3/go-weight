package repository

import (
	"context"

	"github.com/bagasss3/go-weight/internal/model"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type weightRepository struct {
	db *gorm.DB
}

func NewWeighRepository(db *gorm.DB) model.WeightRepository {
	return &weightRepository{
		db: db,
	}
}

func (w *weightRepository) Store(ctx context.Context, weight *model.Weight) error {
	log := logger.WithFields(logger.Fields{
		"ctx":    ctx,
		"weight": weight,
	})
	if err := w.db.WithContext(ctx).Create(weight).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (w *weightRepository) GetAllWeights(ctx context.Context) ([]*model.Weight, error) {
	log := logger.WithFields(logger.Fields{
		"ctx": ctx,
	})

	var weights []*model.Weight
	if err := w.db.WithContext(ctx).Order("created_at desc").Find(&weights).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return weights, nil
}

func (w *weightRepository) GetWeightByID(ctx context.Context, id int64) (*model.Weight, error) {
	log := logger.WithFields(logger.Fields{
		"ctx": ctx,
		"id":  id,
	})

	var weight *model.Weight
	err := w.db.WithContext(ctx).Take(&weight, "id = ?", id).Error

	switch err {
	case nil:
		return weight, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		log.Error(err)
		return nil, err
	}

}

func (w *weightRepository) UpdateWeight(ctx context.Context, id int64, weight *model.Weight) error {
	log := logger.WithFields(logger.Fields{
		"ctx":    ctx,
		"id":     id,
		"weight": weight,
	})

	if err := w.db.WithContext(ctx).Where("id = ?", id).Updates(weight).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (w *weightRepository) DeleteWeight(ctx context.Context, id int64) (bool, error) {
	log := logger.WithFields(logger.Fields{
		"ctx": ctx,
		"id":  id,
	})

	var weight *model.Weight
	if err := w.db.WithContext(ctx).Where("id = ?", id).Delete(&weight).Error; err != nil {
		log.Error(err)
		return false, err
	}
	return true, nil
}
