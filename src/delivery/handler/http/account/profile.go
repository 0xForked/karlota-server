package account

import (
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
