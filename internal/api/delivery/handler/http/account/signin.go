package account

import (
	"github.com/aasumitro/karlota/internal/api/domain"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// login godoc
// @Schemes
// @Summary Logged User In
// @Description Generate Access Token (JWT).
// @Tags AccountHandler
// @Accept mpfd
// @Produce json
// @Param email formData string true "email address"
// @Param password formData string true "password"
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/login [POST]
func (handler *accountHandler) signIn(context *gin.Context) {
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
