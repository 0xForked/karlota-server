package mysql_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/karlota/internal/api/domain"
	mysql2 "github.com/aasumitro/karlota/internal/api/repository/mysql"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type accountRepositoryTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	accountRepository mysql2.AccountRepository
	user              domain.User
}

// SetupSuite is useful in cases where the setup code is time-consuming and isn't modified in any of the tests.
// An example of when this could be useful is if you were testing code that reads from a database,
// and all the tests used the same data and only ran SELECT statements. In this scenario,
// SetupSuite could be used once to load the database with data.
func (suite *accountRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, suite.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(suite.T(), err)

	suite.accountRepository = mysql2.AccountRepositoryImpl(suite.DB)
}

func (suite *accountRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_Find() {
	password := utils.Hash{}.Make("password")
	suite.user = domain.User{
		ID:       1,
		Name:     "test name",
		Email:    "test@email.com",
		Password: password,
	}
	user := suite.mock.NewRows([]string{"id", "name", "email", "password"})
	user.AddRow(suite.user.ID, suite.user.Name, suite.user.Email, suite.user.Password)
	//if you have expression (i.e. use UPDATE or INSERT), you should use ExpectExec
	//if you have query (i.e. use SELECT), you should use ExpectQuery
	suite.mock.ExpectQuery("SELECT").WithArgs(suite.user.Email).WillReturnRows(user)
	res, err := suite.accountRepository.Find(suite.user.Email)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
	require.NoError(suite.T(), err)
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_Find_NotFound() {
	suite.mock.ExpectQuery("SELECT").
		WithArgs("test@email.wrong").
		WillReturnError(errors.New("NOT_FOUND"))
	res, err := suite.accountRepository.Find("test@email.wrong")
	require.Equal(suite.T(), err.Error(), "NOT_FOUND")
	require.Equal(suite.T(), &domain.User{}, res)
	require.Error(suite.T(), err)
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_Store() {
	password := utils.Hash{}.Make("password")
	suite.user = domain.User{
		Name:     "test name",
		Email:    "test@email.com",
		FCMToken: "1234",
		IsOnline: false,
		Password: password,
	}
	suite.mock.MatchExpectationsInOrder(false)
	suite.mock.ExpectBegin()
	//if you have expression (i.e. use UPDATE or INSERT), you should use ExpectExec
	//if you have query (i.e. use SELECT), you should use ExpectQuery
	//query := `insert into test_table \(name, email, password\) values \(\$1, \$2\,\$3\)`
	suite.mock.ExpectExec("INSERT").
		WithArgs(
			suite.user.Name,
			suite.user.Email,
			suite.user.FCMToken,
			suite.user.IsOnline,
			suite.user.Password,
		).
		WillReturnResult(driver.ResultNoRows).
		WillReturnError(nil)
	suite.mock.ExpectCommit()
	err := suite.accountRepository.Store(&suite.user)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_Store_Error() {
	password := utils.Hash{}.Make("password")
	suite.user = domain.User{
		Name:     "test name",
		Email:    "test@email.com",
		FCMToken: "1234",
		IsOnline: false,
		Password: password,
	}
	suite.mock.MatchExpectationsInOrder(false)
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("INSERT").
		WithArgs(
			suite.user.Name,
			suite.user.Email,
			suite.user.FCMToken,
			suite.user.IsOnline,
			suite.user.Password,
		).
		WillReturnError(errors.New("FAILED_SOMETHING_WENT_WRONG"))
	suite.mock.ExpectRollback()
	err := suite.accountRepository.Store(&suite.user)
	require.NotNil(suite.T(), err)
	require.Equal(suite.T(), "FAILED_SOMETHING_WENT_WRONG", err.Error())
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_All_ShouldSuccess() {
	user := suite.mock.
		NewRows([]string{"id", "name", "email"}).
		AddRow(1, "test name", "test@email.com").
		AddRow(2, "test name 2", "test2@email.com")

	suite.mock.
		ExpectQuery("SELECT").
		WithArgs(). // TODO IF NEED PAGINATION
		WillReturnRows(user)

	res, err := suite.accountRepository.All()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
	require.NoError(suite.T(), err)
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_All_ShouldError() {
	suite.mock.ExpectQuery("SELECT").
		WillReturnError(errors.New("NOT_FOUND"))
	_, err := suite.accountRepository.All()
	require.Equal(suite.T(), err.Error(), "NOT_FOUND")
	require.Error(suite.T(), err)
}

func TestAccountRepository(t *testing.T) {
	suite.Run(t, new(accountRepositoryTestSuite))
}
