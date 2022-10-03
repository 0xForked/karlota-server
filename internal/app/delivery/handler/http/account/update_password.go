package account

import (
	"github.com/aasumitro/karlota/internal/app/domain"
	utils2 "github.com/aasumitro/karlota/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// update password godoc
// @Schemes
// @Summary Update Password
// @Description Generate New Password.
// @Tags AccountHandler
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param password formData string true "new password"
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/update/password [POST]
func (handler *accountHandler) updatePassword(context *gin.Context) {
	var form domain.UserPasswordForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils2.NewFormRequest(domain.UserFormErrorMessages).Validate(form, err)
		utils2.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	payload := context.MustGet("payload")
	email := payload.(map[string]interface{})["email"]
	if err := handler.service.Edit(&domain.User{
		Email:    email.(string),
		Password: form.Password,
	}); err != nil {
		utils2.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	utils2.NewHttpRespond(context, http.StatusOK, "UPDATED")
}
