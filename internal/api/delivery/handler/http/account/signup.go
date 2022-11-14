package account

import (
	"github.com/aasumitro/karlota/internal/api/domain"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// register godoc
// @Schemes
// @Summary Register new User
// @Description Generate new User Account.
// @Tags AccountHandler
// @Accept mpfd
// @Produce json
// @Param name formData string true "full name"
// @Param email formData string true "email address"
// @Param password formData string true "password"
// @Success 201 {object} utils.SuccessRespond "CREATED_RESPOND"
// @Failure 400 {object} utils.ErrorRespond "BAD_REQUEST_RESPOND"
// @Failure 422 {object} utils.ValidationErrorRespond "UNPROCESSABLE_ENTITY_RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL_SERVER_ERROR_RESPOND"
// @Router /v1/register [POST]
func (handler *accountHandler) signUp(context *gin.Context) {
	var form domain.UserRegisterForm

	if err := context.ShouldBind(&form); err != nil {
		validationError := utils.NewFormRequest(domain.UserFormErrorMessages).Validate(form, err)
		utils.NewHttpRespond(context, http.StatusUnprocessableEntity, validationError)
		return
	}

	newUser := domain.User{Name: form.Name, Email: form.Email, Password: form.Password}
	if err := handler.service.Register(&newUser); err != nil {
		utils.NewHttpRespond(context, http.StatusBadRequest, err.Error())
		return
	}

	utils.NewHttpRespond(context, http.StatusCreated, "ACCOUNT_CREATED")
}
