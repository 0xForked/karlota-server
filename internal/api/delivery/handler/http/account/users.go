package account

import (
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// users godoc
// @Schemes
// @Summary User List
// @Description Get User List.
// @Tags AccountHandler
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} utils.SuccessRespond "OK_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/users [GET]
func (handler *accountHandler) users(context *gin.Context) {
	users, err := handler.service.List()
	if err != nil {
		utils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
	}

	utils.NewHttpRespond(context, http.StatusOK, users)
}
