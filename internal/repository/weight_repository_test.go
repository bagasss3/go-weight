package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagasss3/go-weight/internal/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeWeightRepositoryWithMock(db *gorm.DB) model.WeightRepository {
	return NewWeighRepository(db)
}

func initializeDBMockConn() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func TestWeightRepository_Store(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectQuery(`INSERT INTO "weights"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
				sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		dbmock.ExpectCommit()

		err := weightRepo.Store(ctx, &model.Weight{
			ID: 1,
		})
		assert.NoError(t, err)
	})
	t.Run("failed", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectQuery(`INSERT INTO "weights"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
				sqlmock.AnyArg()).
			WillReturnError(gorm.ErrInvalidDB)
		dbmock.ExpectRollback()

		err := weightRepo.Store(ctx, &model.Weight{
			ID: 1,
		})
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrInvalidDB, err)
	})
}

func TestWeightRepository_GetAllWeights(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		res := sqlmock.NewRows([]string{"id"}).AddRow(1)
		dbmock.ExpectQuery("^SELECT .+ FROM \"weights\"").WillReturnRows(res)

		weights, err := weightRepo.GetAllWeights(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(weights))
	})
	t.Run("failed", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()

		dbmock.ExpectQuery("^SELECT .+ FROM \"weights\"").WillReturnError(gorm.ErrInvalidDB)

		weights, err := weightRepo.GetAllWeights(ctx)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrInvalidDB, err)
		assert.Nil(t, weights)
	})

}

func TestWeightRepository_GetWeightByID(t *testing.T) {
	weightMock := &model.Weight{
		ID: 1,
	}
	t.Run("success", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		res := sqlmock.NewRows([]string{"id"}).AddRow(weightMock.ID)
		dbmock.ExpectQuery("^SELECT .+ FROM \"weights\"").WillReturnRows(res)

		weights, err := weightRepo.GetWeightByID(ctx, weightMock.ID)
		assert.NoError(t, err)
		assert.Equal(t, weights, weightMock)
	})
	t.Run("failed", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()

		dbmock.ExpectQuery("^SELECT .+ FROM \"weights\"").WillReturnError(gorm.ErrInvalidDB)

		weights, err := weightRepo.GetWeightByID(ctx, weightMock.ID)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrInvalidDB, err)
		assert.Nil(t, weights)
	})

}

func TestWeightRepository_UpdateWeight(t *testing.T) {
	weightMock := &model.Weight{
		ID:         1,
		MaxWeight:  5,
		MinWeight:  4,
		Difference: 1,
		CreatedAt:  time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}
	t.Run("success", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectExec(`UPDATE "weights"`).
			WithArgs(weightMock.MaxWeight, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), weightMock.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		dbmock.ExpectCommit()

		err := weightRepo.UpdateWeight(ctx, weightMock.ID, weightMock)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectExec(`UPDATE "weights"`).
			WithArgs(weightMock.MaxWeight, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), weightMock.ID).
			WillReturnError(gorm.ErrInvalidDB)
		dbmock.ExpectRollback()

		err := weightRepo.UpdateWeight(ctx, weightMock.ID, weightMock)
		assert.Error(t, err)
	})
}

func TestWeightRepository_DeleteWeight(t *testing.T) {
	weightMock := &model.Weight{
		ID: 1,
	}
	t.Run("success", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectExec("^DELETE FROM \"weights\"").
			WithArgs(weightMock.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		dbmock.ExpectCommit()

		res, err := weightRepo.DeleteWeight(ctx, weightMock.ID)
		assert.NoError(t, err)
		assert.Equal(t, true, res)
	})

	t.Run("failed", func(t *testing.T) {
		db, dbmock := initializeDBMockConn()
		weightRepo := initializeWeightRepositoryWithMock(db)
		ctx := context.TODO()
		dbmock.ExpectBegin()
		dbmock.ExpectExec("^DELETE FROM \"weights\"").
			WithArgs(weightMock.ID).
			WillReturnError(gorm.ErrInvalidDB)
		dbmock.ExpectRollback()

		res, err := weightRepo.DeleteWeight(ctx, weightMock.ID)
		assert.Error(t, err)
		assert.Equal(t, false, res)
	})
}
