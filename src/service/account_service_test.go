package service_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/karlota/src/domain"
	"github.com/aasumitro/karlota/src/domain/mocks"
	"github.com/aasumitro/karlota/src/service"
	"github.com/aasumitro/karlota/src/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type accountTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	user  domain.User
	users []domain.User
	jwt   utils.JWT
}

func (suite *accountTestSuite) SetupTest() {
	suite.jwt = utils.JWT{
		SecretKey:       "secret$%^!@12345://@()",
		ExpirationHours: 24,
		Issuer:          "KARLOTA_TEST",
	}

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
	accountRepository := new(mocks.AccountRepository)
	suite.user.Password = "password"
	accountRepository.
		On("Store", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	err := svc.Register(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin() {
	accountRepository := new(mocks.AccountRepository)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	_, err := svc.Login(suite.user.Email, "password")
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_NotFoundEmail() {
	accountRepository := new(mocks.AccountRepository)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(nil, errors.New("EMAIL_NOT_FOUND")).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	_, err := svc.Login("test@email.wrong", "password")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "EMAIL_NOT_FOUND")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_InvalidPassword() {
	accountRepository := new(mocks.AccountRepository)
	suite.user.Password = utils.Hash{}.Make("password")
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	_, err := svc.Login(suite.user.Email, "123456")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "INVALID_PASSWORD")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestProfile() {
	accountRepository := new(mocks.AccountRepository)
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	user, err := svc.Profile(mock.Anything)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), user, &suite.user)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestListAccount() {
	accountRepository := new(mocks.AccountRepository)
	accountRepository.
		On("All", mock.Anything).
		Return(&suite.users, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	users, err := svc.List()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), users, &suite.users)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdateFCMToken() {
	accountRepository := new(mocks.AccountRepository)
	suite.user.FCMToken = "loremIpsum-12345"
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	err := svc.Edit(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestUpdatePassword() {
	accountRepository := new(mocks.AccountRepository)
	suite.user.Password = "lorem"
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	accountRepository.
		On("Edit", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository, suite.jwt)
	err := svc.Edit(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
