package httpsvc

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bagasss3/go-weight/internal/model"
	"github.com/bagasss3/go-weight/internal/model/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHTTP_handleCreateWeight(t *testing.T) {
	weight := &model.Weight{
		ID:        1,
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeight := mock.NewMockWeightController(ctrl)
	ec := echo.New()

	server := &Service{
		group:            ec.Group("/api"),
		weightController: mockWeight,
	}

	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/weight", strings.NewReader(`
        {
        "max_weight": 10,
        "min_weight": 5,
        "created_at": "2020-05-01"
        }`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		// maxWeight := ectx.FormValue("")
		mockWeight.EXPECT().Create(ectx.Request().Context(), &model.WeightInput{
			MaxWeight: 10,
			MinWeight: 5,
			CreatedAt: "2020-05-01",
		}).Times(1).Return(weight, nil)

		err := server.handleCreateWeight()(ectx)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("failed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/weight", strings.NewReader(`
        {
        "max_weight": 10,
        "min_weight": 5,
        "created_at": "2020-05-01"
        }`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)

		mockWeight.EXPECT().Create(ectx.Request().Context(), &model.WeightInput{
			MaxWeight: 10,
			MinWeight: 5,
			CreatedAt: "2020-05-01",
		}).Times(1).Return(nil, errors.New("usecase err"))

		err := server.handleCreateWeight()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.Error(t, err)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func TestHTTP_handleGetAllWeights(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeight := mock.NewMockWeightController(ctrl)
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
	ec := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/weight", nil)

	server := &Service{
		group:            ec.Group("/api"),
		weightController: mockWeight,
	}

	t.Run("ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		mockWeight.EXPECT().ReadWeights(ectx.Request().Context()).Times(1).Return(weights, nil)

		err := server.handleGetAllWeights()(ectx)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("failed", func(t *testing.T) {
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		mockWeight.EXPECT().ReadWeights(ectx.Request().Context()).Times(1).Return(nil, errors.New("usecase err"))

		err := server.handleGetAllWeights()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.Error(t, err)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func TestHTTP_handleShowWeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeight := mock.NewMockWeightController(ctrl)
	weight := &model.Weight{
		ID:        1,
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}
	ec := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/weight", nil)

	server := &Service{
		group:            ec.Group("/api"),
		weightController: mockWeight,
	}

	t.Run("ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")
		mockWeight.EXPECT().ShowWeight(ectx.Request().Context(), weight.ID).Times(1).Return(weight, nil)

		err := server.handleShowWeight()(ectx)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("failed", func(t *testing.T) {
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")
		mockWeight.EXPECT().ShowWeight(ectx.Request().Context(), weight.ID).Times(1).Return(nil, errors.New("usecase err"))

		err := server.handleShowWeight()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.Error(t, err)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func TestHTTP_handleUpdateWeight(t *testing.T) {
	weight := &model.Weight{
		ID:        1,
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeight := mock.NewMockWeightController(ctrl)
	ec := echo.New()

	server := &Service{
		group:            ec.Group("/api"),
		weightController: mockWeight,
	}

	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/weight", strings.NewReader(`
        {
        "max_weight": 10,
        "min_weight": 5,
        "created_at": "2020-05-01"
        }`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")

		mockWeight.EXPECT().UpdateWeight(ectx.Request().Context(), weight.ID, &model.WeightInput{
			MaxWeight: 10,
			MinWeight: 5,
			CreatedAt: "2020-05-01",
		}).Times(1).Return(weight, nil)

		err := server.handleUpdateWeight()(ectx)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("failed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/api/weight", strings.NewReader(`
        {
        "max_weight": 10,
        "min_weight": 5,
        "created_at": "2020-05-01"
        }`,
		))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")

		mockWeight.EXPECT().UpdateWeight(ectx.Request().Context(), weight.ID, &model.WeightInput{
			MaxWeight: 10,
			MinWeight: 5,
			CreatedAt: "2020-05-01",
		}).Times(1).Return(nil, errors.New("controller err"))

		err := server.handleUpdateWeight()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.Error(t, err)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func TestHTTP_handleDeleteWeight(t *testing.T) {
	weight := &model.Weight{
		ID:        1,
		MaxWeight: 10,
		MinWeight: 5,
		CreatedAt: time.Date(2020, time.May, 01, 00, 00, 00, 00, time.UTC),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeight := mock.NewMockWeightController(ctrl)
	ec := echo.New()

	server := &Service{
		group:            ec.Group("/api"),
		weightController: mockWeight,
	}

	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/weight", nil)
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")

		mockWeight.EXPECT().DeleteWeight(ectx.Request().Context(), weight.ID).Times(1).Return(true, nil)

		err := server.handleDeleteWeight()(ectx)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, rec.Result().StatusCode)
	})

	t.Run("failed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/weight", nil)
		rec := httptest.NewRecorder()
		ectx := ec.NewContext(req, rec)
		ectx.SetParamNames("id")
		ectx.SetParamValues("1")

		mockWeight.EXPECT().DeleteWeight(ectx.Request().Context(), weight.ID).Times(1).Return(false, errors.New("controller err"))

		err := server.handleDeleteWeight()(ectx)
		ec.DefaultHTTPErrorHandler(err, ectx)
		require.Error(t, err)
		require.EqualValues(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}
