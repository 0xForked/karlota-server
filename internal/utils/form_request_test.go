package utils_test

import (
	"bytes"
	"encoding/json"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ! MOCK  ======================================================================================

type mockObject struct {
	Email    string `json:"email" binding:"required,email" msg:"error_invalid_email"`
	Username string `json:"username" binding:"required,alphanum,gte=6,lte=32" msg:"error_invalid_username"`
	Password string `json:"password" binding:"required,gte=6,lte=32" msg:"error_invalid_password"`
}

type mockObjectNoTag struct {
	TestNoTag string `json:"test_no_tag" binding:"required,gte=6,lte=32" msg:"error_test_not_tag"`
}

var mockMessage = map[string]string{
	"error_invalid_email":    "please enter a valid email address",
	"error_invalid_username": "username must be alphanumeric and between 6 and 32 characters",
	"error_invalid_password": "password must be between 6 and 32 characters",
}

func mockServer(context *gin.Context, content interface{}) {
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

// ! MOCK  ======================================================================================

type formRequestTestSuite struct {
	suite.Suite
	context     *gin.Context
	formRequest utils.FormRequest
}

func (suite *formRequestTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.formRequest = utils.NewFormRequest(mockMessage)
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.context.Request = &http.Request{Header: make(http.Header)}
}

func (suite *formRequestTestSuite) TestRequestInvalidAllForm() {
	mockServer(suite.context, map[string]interface{}{
		"username": "test123!@#%",
		"email":    "testemail.com",
		"password": "12345",
	})

	req := mockObject{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.formRequest.Validate(req, err)
		assert.True(suite.T(), len(errors) == 3)

		valueUsername, ok := errors["username"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valueUsername == mockMessage["error_invalid_username"])

		valueEmail, ok := errors["email"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valueEmail == mockMessage["error_invalid_email"])

		valuePassword, ok := errors["password"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valuePassword == mockMessage["error_invalid_password"])
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

func (suite *formRequestTestSuite) TestRequestInvalidAllForm_NoTag() {
	mockServer(suite.context, map[string]interface{}{
		"test_no_tag": "12345",
	})

	req := mockObjectNoTag{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.formRequest.Validate(req, err)
		assert.True(suite.T(), len(errors) == 1)

		valueTestNoTag, ok := errors["test_no_tag"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valueTestNoTag == "error_test_not_tag")
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

func TestFormRequest(t *testing.T) {
	suite.Run(t, new(formRequestTestSuite))
}
