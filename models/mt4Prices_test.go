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

type Mt4Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	mt4       Mt4
	mt4_price *Mt4Price
}

func (m *Mt4Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, m.mock, err = sqlmock.New()
	require.NoError(m.T(), err)

	m.DB, err = gorm.Open("mysql", db)
	require.NoError(m.T(), err)

	//s.DB.LogMode(true)

	m.mt4 = CreateMt4Prices(m.DB)
}

func (m *Mt4Suite) AfterTest(_, _ string) {
	require.NoError(m.T(), m.mock.ExpectationsWereMet())
}

func TestMt4Init(t *testing.T) {
	suite.Run(t, new(Mt4Suite))
}

func (m *Mt4Suite) TestGetExchangeRatio() {
	var (
		testBid = 0.66
		testAsk = 0.66
	)
	selectQuery := fmt.Sprintf("SELECT bid,ask FROM `mt4_prices`  WHERE \\(symbol='AUDUSD'\\) ORDER BY time desc LIMIT 1")
	m.mock.ExpectQuery(selectQuery).
		WillReturnRows(sqlmock.NewRows([]string{"bid", "ask"}).
			AddRow(testBid, testAsk))
	bid, ask, err := m.mt4.GetExchangeRatio()
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(m.T(), err)
	require.Nil(m.T(), deep.Equal(testBid, bid))
	require.Nil(m.T(), deep.Equal(testAsk, ask))
}
