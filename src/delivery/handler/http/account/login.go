package account

import (
	"github.com/aasumitro/karlota/src/domain"
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// login handle.
// @Schemes
// @Summary Logged User In
// @Description Generate Access Token (JWT).
// @Tags Karlota Messaging
// @Accept mpfd
// @Produce json
// @Param email formData string true "email address"
// @Param password formData string true "password"
// @Success 201 {object} delivery.HttpSuccessRespond{data=object} "CREATED_RESPOND"
// @Failure 400 {object} delivery.HttpErrorRespond{data=string} "BAD_REQUEST_RESPOND"
// @Failure 422 {object} delivery.HttpValidationErrorRespond{data=object} "UNPROCESSABLE_ENTITY_RESPOND"
// @Router /v1/login [POST]
func (handler *accountHandler) login(context *gin.Context) {
	var form domain.UserLoginForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewFormRequest(domain.UserFormErrorMessages).Validate(form, err)
		utils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	payload, err := handler.service.Login(form.Email, form.Password)
	if err != nil {
		utils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	utils.NewHttpRespond(context, http.StatusCreated, payload)
}
