package models

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CashSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	cash             Cash
	cash_requirement *CashRequirement
}

func (c *CashSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, c.mock, err = sqlmock.New()
	require.NoError(c.T(), err)

	c.DB, err = gorm.Open("mysql", db)
	require.NoError(c.T(), err)

	//s.DB.LogMode(true)

	c.cash = CreateCashRequirement(c.DB)
}

func (c *CashSuite) AfterTest(_, _ string) {
	require.NoError(c.T(), c.mock.ExpectationsWereMet())
}

func TestCashInit(t *testing.T) {
	suite.Run(t, new(CashSuite))
}

func (c *CashSuite) TestGetCashRequirement() {
	var (
		id              = "1"
		cashrequirement = 100.12
		maxvalueatrisk  = 200.23
		primebrokercash = 300.34
		totalcash       = 400.45
	)
	selectQuery := fmt.Sprintf("SELECT id,cash_requirement,max_value_at_risk,prime_broker_cash,total_cash FROM `cash_requirement`  ORDER BY id desc LIMIT 1")
	c.mock.ExpectQuery(selectQuery).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cashrequirement", "maxvalueatrisk", "primebrokercash", "totalcash"}).
			AddRow(id, cashrequirement, maxvalueatrisk, primebrokercash, totalcash))
	res, err := c.cash.GetCashRequirement()
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(c.T(), err)
	require.Nil(c.T(), deep.Equal(CashRequirement{ID: id, CashRequirement: cashrequirement, MaxValueAtRisk: maxvalueatrisk, PrimeBrokerCash: primebrokercash, TotalCash: totalcash}, res))
}
