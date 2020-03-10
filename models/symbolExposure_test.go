package models

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SymbolSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	sym          Sym
	sym_exposure *SymExposure
}

func (s *SymbolSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	//s.DB.LogMode(true)

	s.sym = CreateSymExposure(s.DB)
}

func (s *SymbolSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestSymInit(t *testing.T) {
	suite.Run(t, new(SymbolSuite))
}

func (s *SymbolSuite) TestGetSymExposure() {
	var (
		id          = "1"
		timestamp   = "2020-02-24 10:10:10"
		sym         = "AUDUSD"
		cptype      = "Client"
		daypnl      = -200.455
		exposure    = -100.3342
		realisedpnl = -100.123
	)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT id,ts,sym,cptype,exposure,daypnl,realisedpnl FROM `sym_exposure` WHERE (ts = (SELECT MAX(ts) FROM `sym_exposure` )) LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "timestamp", "sym", "cptype", "exposure", "daypnl", "realisedpnl"}).
			AddRow(id, timestamp, sym, cptype, exposure, daypnl, realisedpnl))
	res, err := s.sym.GetSymExposure()
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(SymExposure{ID: id, Timestamp: timestamp, Sym: sym, CpType: cptype, Exposure: exposure, DayPnl: daypnl, RealisedPnl: realisedpnl}, res))
}

func (s *SymbolSuite) TestGetSumRealisedPnl() {
	var (
		sum = 500.0
	)
	sumQuery := fmt.Sprintf("SELECT Sum\\(realisedpnl\\) FROM `sym_exposure` WHERE \\(ts \\= \\(SELECT MAX\\(ts\\) FROM `sym_exposure` \\)\\)")
	s.mock.ExpectQuery(sumQuery).
		WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(sum))

	TestSum, err := s.sym.GetSumRealisedPnl()

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(TestSum, sum))

}
