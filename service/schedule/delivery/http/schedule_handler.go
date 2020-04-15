package http

import (
	"github.com/labstack/echo"
	"github.com/models"
	"github.com/service/promo"
	"github.com/service/schedule"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// promoHandler  represent the httphandler for promo
type scheduleHandler struct {
	promoUsecase promo.Usecase
	scheduleUsecase schedule.Usecase
}

// NewpromoHandler will initialize the promos/ resources endpoint
func NewScheduleHandler(e *echo.Echo, us promo.Usecase,su schedule.Usecase) {
	handler := &scheduleHandler{
		promoUsecase: us,
		scheduleUsecase:su,
	}
	//e.POST("/promos", handler.Createpromo)
	//e.PUT("/promos/:id", handler.Updatepromo)
	e.GET("service/schedule", handler.GetSchedule)
	//e.GET("service/special-promo/:code", handler.GetPromoByCode)
	//e.DELETE("/promos/:id", handler.Delete)
}

func isRequestValid(m *models.NewCommandMerchant) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetByID will get article by given id
func (a *scheduleHandler) GetSchedule(c echo.Context) error {
	merchantId := c.QueryParam("merchant_id")
	//qsize := c.QueryParam("size")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

		art, err := a.scheduleUsecase.GetScheduleByMerchantId(ctx, merchantId)
		if err != nil {
			return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, art)

}


func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrUnAuthorize:
		return http.StatusUnauthorized
	case models.ErrConflict:
		return http.StatusBadRequest
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
