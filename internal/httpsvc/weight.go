package httpsvc

import (
	"net/http"
	"strconv"

	"github.com/bagasss3/go-weight/internal/model"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

func (s *Service) handleCreateWeight() echo.HandlerFunc {
	return func(c echo.Context) error {
		var body *model.WeightInput
		if err := c.Bind(&body); err != nil {
			logger.Error(err)
			return err
		}

		weightNew, err := s.weightController.Create(c.Request().Context(), body)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "something's wrong")
		}
		return c.JSON(http.StatusOK, weightNew)
	}
}

func (s *Service) handleGetAllWeights() echo.HandlerFunc {
	return func(c echo.Context) error {
		weights, err := s.weightController.ReadWeights(c.Request().Context())
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "something's wrong")
		}
		return c.JSON(http.StatusOK, weights)

	}
}

func (s *Service) handleShowWeight() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.NoContent(http.StatusPreconditionFailed)
		}
		weight, err := s.weightController.ShowWeight(c.Request().Context(), id)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "something's wrong")
		}
		if weight == nil {
			return echo.NewHTTPError(http.StatusNotFound, "record not found")
		}

		return c.JSON(http.StatusOK, weight)
	}
}

func (s *Service) handleUpdateWeight() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.NoContent(http.StatusPreconditionFailed)
		}

		var body *model.WeightInput
		if err := c.Bind(&body); err != nil {
			logger.Error(err)
			return err
		}

		updatedWeight, err := s.weightController.UpdateWeight(c.Request().Context(), id, body)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "something's wrong")
		}
		return c.JSON(http.StatusOK, updatedWeight)
	}
}

func (s *Service) handleDeleteWeight() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.NoContent(http.StatusPreconditionFailed)
		}

		res, err := s.weightController.DeleteWeight(c.Request().Context(), id)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "something's wrong")
		}
		return c.JSON(http.StatusOK, res)
	}
}
