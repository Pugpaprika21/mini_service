package user

import (
	"miniservice/app/internal/domain/dto/qryparam"
	"miniservice/app/internal/domain/dto/request"
	"miniservice/app/internal/enum"
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
		return c.JSON(http.StatusBadRequest, resp.Message("Error Body With: "+err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build())
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message("Body validation error: "+err.Error()).Code(enum.FOR_BAD_REQUEST).RequestID(reqID).Build())
	}

	data, totalRow, err := u.service.GetUsers(c.Request().Context(), &req, &qry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp.Message("Internal Server Error with: "+err.Error()).Code(enum.FOR_ERROR).RequestID(reqID).Build())
	}

	return c.JSON(http.StatusOK, resp.Message(enum.SUCCESS_STR).Code(200).RequestID(reqID).Total(totalRow).Data(data).Build())
}
