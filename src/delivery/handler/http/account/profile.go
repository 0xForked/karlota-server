package account

import (
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// profile godoc
// @Schemes
// @Summary User Profile
// @Description Get User Data in Detail.
// @Tags AccountHandler
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} utils.SuccessRespond "OK_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 401 {object} utils.ErrorRespond "UNAUTHORIZED_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/profile [GET]
func (handler *accountHandler) profile(context *gin.Context) {
	payload := context.MustGet("payload").(interface{})
	email := payload.(map[string]interface{})["email"]

	profile, err := handler.service.Profile(email.(string))
	if err != nil {
		utils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	utils.NewHttpRespond(context, http.StatusOK, profile)
}
