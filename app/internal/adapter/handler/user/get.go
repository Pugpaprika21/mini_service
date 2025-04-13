package user

import (
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u *userhandler) GetUsers(c echo.Context) error {
	var req request.GetUsers
	var qry qryparam.GetUsers
	var reqID = c.Response().Header().Get(echo.HeaderXRequestID)
	var resp = response.NewResponseBuilder()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message("Error Body With: "+err.Error()).Code(400).Build())
	}

	if err := c.Bind(&qry); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message("Error Query String With: "+err.Error()).Code(400).Build())
	}

	data, err := u.service.GetUsers(c.Request().Context(), &req, &qry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message("Internal Server Error with: "+err.Error()).Code(500).Build())
	}

	return c.JSON(http.StatusOK, resp.Message("Success").Code(200).Data(data).RequestID(reqID).Build())
}
