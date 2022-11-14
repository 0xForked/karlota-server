package service_test

import (
	"errors"
	"github.com/aasumitro/karlota/internal/api/domain"
	domainMocks "github.com/aasumitro/karlota/internal/api/domain/mocks"
	"github.com/aasumitro/karlota/internal/api/service"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type accountTestSuite struct {
	suite.Suite
	DB *gorm.DB

	user  domain.User
	users []domain.User
}

func (suite *accountTestSuite) SetupTest() {
	suite.user = domain.User{
		ID:    1,
		Name:  "Test User",
		Email: "user@test.id",
	}

	suite.users = []domain.User{
		{
			ID:    1,
			Name:  "Test User",
			Email: "user@test.id",
		},
		{
			ID:    2,
			Name:  "Test User Two",
			Email: "user2@test.id",
		},
	}
}

func (suite *accountTestSuite) TestRegister() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = "password"
	accountRepository.
		On("Store", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	err := svc.Register(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	jsonWebTokenUtil.
		On("GetExpirationHours", mock.Anything).
		Return(24)
	jsonWebTokenUtil.
		On("Claim", mock.Anything).
		Return("lorem", nil)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	_, err := svc.Login(suite.user.Email, "password")
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_NotFoundEmail() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(nil, errors.New("EMAIL_NOT_FOUND")).
		Once()
	jsonWebTokenUtil.
		On("GetExpirationHours", mock.Anything).
		Return(24)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	_, err := svc.Login("test@email.wrong", "password")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "EMAIL_NOT_FOUND")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_InvalidPassword() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	jsonWebTokenUtil.
		On("GetExpirationHours", mock.Anything).
		Return(24)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	_, err := svc.Login(suite.user.Email, "123456")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "INVALID_PASSWORD")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_FailedClaimJWT() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	jsonWebTokenUtil.
		On("GetExpirationHours", mock.Anything).
		Return(24)
	jsonWebTokenUtil.
		On("Claim", mock.Anything).
		Return("", errors.New("FAILED_GENERATE_JWT"))
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	_, err := svc.Login(suite.user.Email, "password")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "FAILED_GENERATE_JWT")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestProfile() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	user, err := svc.Profile(mock.Anything)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), user, &suite.user)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestListAccount() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	accountRepository.
		On("All", mock.Anything).
		Return(&suite.users, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	users, err := svc.List()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), users, &suite.users)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdateFCMToken() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.FCMToken = "loremIpsum-12345"
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	accountRepository.
		On("Update", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	err := svc.Edit(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdatePassword() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	suite.user.Password = "lorem"
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	accountRepository.
		On("Update", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	err := svc.Edit(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdate_ShouldError_USER_NOT_FOUND() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	accountRepository.
		On("Find", mock.Anything).
		Return(nil, errors.New("USER_NOT_FOUND")).
		Once()
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	err := svc.Edit(&suite.user)
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "USER_NOT_FOUND")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdate_ShouldError_FAILED_UPDATE_DATA() {
	accountRepository := new(domainMocks.AccountRepository)
	jsonWebTokenUtil := new(domainMocks.JSONWebToken)
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	accountRepository.
		On("Update", &suite.user).
		Return(errors.New("FAILED_UPDATE_DATA"))
	svc := service.AccountServiceImpl(accountRepository, jsonWebTokenUtil)
	err := svc.Edit(&suite.user)
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "FAILED_UPDATE_DATA")
	accountRepository.AssertExpectations(suite.T())
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
