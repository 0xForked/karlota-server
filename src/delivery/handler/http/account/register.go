package account

import (
	"github.com/aasumitro/karlota/src/domain"
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *accountHandler) register(context *gin.Context) {
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
