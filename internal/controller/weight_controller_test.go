package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bagasss3/go-weight/internal/model"
	"github.com/bagasss3/go-weight/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWeightController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockWeightRepo := mock.NewMockWeightRepository(ctrl)

	controller := &weightController{
		weightRepo: mockWeightRepo,
	}

	weightReq := &model.WeightInput{
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: "2020-05-01",
	}

	weight := &model.Weight{
		MaxWeight:  weightReq.MaxWeight,
		MinWeight:  weightReq.MinWeight,
		Difference: weightReq.MaxWeight - weightReq.MinWeight,
		CreatedAt:  time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}

	t.Run("ok", func(t *testing.T) {
		mockWeightRepo.EXPECT().Store(ctx, weight).Times(1).Return(nil)
		res, err := controller.Create(ctx, weightReq)
		require.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("fail", func(t *testing.T) {
		mockWeightRepo.EXPECT().Store(ctx, weight).Times(1).Return(errors.New("some error"))
		res, err := controller.Create(ctx, weightReq)
		require.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestWeightController_ReadWeights(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockWeightRepo := mock.NewMockWeightRepository(ctrl)

	controller := &weightController{
		weightRepo: mockWeightRepo,
	}

	weights := []*model.Weight{
		{
			ID:         1,
			MaxWeight:  10,
			MinWeight:  5,
			Difference: 5,
			CreatedAt:  time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
		},
		{
			ID:         2,
			MaxWeight:  50,
			MinWeight:  49,
			Difference: 1,
			CreatedAt:  time.Date(2020, time.May, 02, 00, 00, 00, 00, time.UTC),
		},
	}

	t.Run("ok", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetAllWeights(ctx).Times(1).Return(weights, nil)
		res, err := controller.ReadWeights(ctx)
		require.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("fail", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetAllWeights(ctx).Times(1).Return(nil, errors.New("some error"))
		res, err := controller.ReadWeights(ctx)
		require.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestWeightController_ShowWeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockWeightRepo := mock.NewMockWeightRepository(ctrl)

	controller := &weightController{
		weightRepo: mockWeightRepo,
	}

	weight := &model.Weight{
		ID:        1,
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}

	t.Run("ok", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, weight.ID).Times(1).Return(weight, nil)
		res, err := controller.ShowWeight(ctx, weight.ID)
		require.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run("fail", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, weight.ID).Times(1).Return(nil, errors.New("some error"))
		res, err := controller.ShowWeight(ctx, weight.ID)
		require.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestWeightController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockWeightRepo := mock.NewMockWeightRepository(ctrl)

	controller := &weightController{
		weightRepo: mockWeightRepo,
	}

	weightReq := &model.WeightInput{
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: "2020-05-01",
	}

	oldWeight := &model.Weight{
		ID:         1,
		MaxWeight:  5,
		MinWeight:  3,
		Difference: 2,
		CreatedAt:  time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}

	newWeight := &model.Weight{
		ID:         oldWeight.ID,
		MaxWeight:  weightReq.MaxWeight,
		MinWeight:  weightReq.MinWeight,
		Difference: weightReq.MaxWeight - weightReq.MinWeight,
		CreatedAt:  oldWeight.CreatedAt,
	}

	t.Run("ok", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, oldWeight.ID).Times(1).Return(oldWeight, nil)
		mockWeightRepo.EXPECT().UpdateWeight(ctx, oldWeight.ID, newWeight).Times(1).Return(nil)
		res, err := controller.UpdateWeight(ctx, oldWeight.ID, weightReq)
		require.NoError(t, err)
		assert.NotNil(t, res)
		require.Equal(t, newWeight, res)
	})
	t.Run("fail", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, oldWeight.ID).Times(1).Return(oldWeight, nil)
		mockWeightRepo.EXPECT().UpdateWeight(ctx, oldWeight.ID, newWeight).Times(1).Return(errors.New("ada error"))
		res, err := controller.UpdateWeight(ctx, oldWeight.ID, weightReq)
		require.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestWeightController_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	mockWeightRepo := mock.NewMockWeightRepository(ctrl)

	controller := &weightController{
		weightRepo: mockWeightRepo,
	}

	weight := &model.Weight{
		ID:         1,
		MaxWeight:  5,
		MinWeight:  3,
		Difference: 2,
		CreatedAt:  time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}

	t.Run("ok", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, weight.ID).Times(1).Return(weight, nil)
		mockWeightRepo.EXPECT().DeleteWeight(ctx, weight.ID).Times(1).Return(true, nil)
		res, err := controller.DeleteWeight(ctx, weight.ID)
		require.NoError(t, err)
		assert.NotNil(t, res)
		require.Equal(t, true, res)
	})
	t.Run("fail", func(t *testing.T) {
		mockWeightRepo.EXPECT().GetWeightByID(ctx, weight.ID).Times(1).Return(weight, nil)
		mockWeightRepo.EXPECT().DeleteWeight(ctx, weight.ID).Times(1).Return(false, errors.New("ada error"))
		res, err := controller.DeleteWeight(ctx, weight.ID)
		require.Error(t, err)
		require.Equal(t, false, res)
	})
}
