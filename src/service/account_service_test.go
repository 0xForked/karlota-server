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

	user domain.User
}

func (suite *accountTestSuite) TestRegister() {
	accountRepository := new(mocks.AccountRepository)
	suite.user = domain.User{
		ID:       1,
		Name:     "Test User",
		Email:    "user@test.id",
		Password: "password",
	}
	accountRepository.
		On("Store", &suite.user).
		Return(nil)
	svc := service.AccountServiceImpl(accountRepository)
	err := svc.Register(&suite.user)
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin() {
	accountRepository := new(mocks.AccountRepository)
	password, _ := utils.Hash{}.Make("password")
	suite.user = domain.User{
		ID:       1,
		Name:     "Test User",
		Email:    "user@test.id",
		Password: password,
	}
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository)
	_, err := svc.Login(suite.user.Email, "password")
	require.NoError(suite.T(), err)
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_NotFoundEmail() {
	accountRepository := new(mocks.AccountRepository)
	password, _ := utils.Hash{}.Make("password")
	suite.user = domain.User{
		ID:       1,
		Name:     "Test User",
		Email:    "user@test.id",
		Password: password,
	}
	accountRepository.
		On("Find", mock.Anything).
		Return(nil, errors.New("EMAIL_NOT_FOUND")).
		Once()
	svc := service.AccountServiceImpl(accountRepository)
	_, err := svc.Login("test@email.wrong", "password")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "EMAIL_NOT_FOUND")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestLogin_ShouldError_InvalidPassword() {
	accountRepository := new(mocks.AccountRepository)
	password, _ := utils.Hash{}.Make("password")
	suite.user = domain.User{
		ID:       1,
		Name:     "Test User",
		Email:    "user@test.id",
		Password: password,
	}
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository)
	_, err := svc.Login(suite.user.Email, "123456")
	require.Error(suite.T(), err)
	require.Equal(suite.T(), err.Error(), "INVALID_PASSWORD")
	accountRepository.AssertExpectations(suite.T())
}

func (suite *accountTestSuite) TestProfile() {
	accountRepository := new(mocks.AccountRepository)
	suite.user = domain.User{
		ID:    1,
		Name:  "Test User",
		Email: "user@test.id",
	}
	accountRepository.
		On("Find", mock.Anything).
		Return(&suite.user, nil).
		Once()
	svc := service.AccountServiceImpl(accountRepository)
	user, err := svc.Profile(mock.Anything)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), user, &suite.user)
	accountRepository.AssertExpectations(suite.T())
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
