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
	//"gopkg.in/DATA-DOG/go-sqlmock.v1"
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

func TestInit(t *testing.T) {
	suite.Run(t, new(SymbolSuite))
}

func (s *SymbolSuite) TestGetSymExposure() {
	var (
		id          = "1"
		timestamp   = "2020-02-24 10:10:10"
		sym         = "AUDUSD"
		cptype      = "Client"
		daypnl      = (float64)(-200.0)
		exposure    = (float64)(-100.0)
		realisedpnl = (float64)(-100.0)
	)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT id,ts,sym,cptype,exposure,daypnl,realisedpnl FROM `sym_exposure` WHERE (ts = (SELECT MAX(ts) FROM `sym_exposure` )) LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "timestamp", "sym", "cptype", "exposure", "daypnl", "realisedpnl"}).
			AddRow(id, timestamp, sym, cptype, exposure, daypnl, realisedpnl))
	//fmt.Println(regexp.QuoteMeta("SELECT id,ts,sym,cptype,exposure,daypnl,realisedpnl FROM `sym_exposure` WHERE (ts = (SELECT MAX(ts) FROM `sym_exposure` )) LIMIT 1"))
	fmt.Println(0)
	res, err := s.sym.GetSymExposure()
	fmt.Println(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(2)
	fmt.Println("Hello \t:", res)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(SymExposure{ID: id, Timestamp: timestamp, Sym: sym, CpType: cptype, Exposure: exposure, DayPnl: daypnl, RealisedPnl: realisedpnl}, res))
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// func (s *Suite) TestGetSumRealisedPnl() {
// 	var (
// 		sum = 500.0
// 	)

// 	s.mock.ExpectQuery(regexp.QuoteMeta(
// 		`SELECT SUM(realisedpnl) FROM sym_exposure WHERE (ts=(SELECT MAX(ts) FROM sym_exposure))`)).
// 		WillReturnRows(sqlmock.NewRows([]string{"SUM(realisedpnl)"}).AddRow(sum))

// 	TestSum, err := s.sym.GetSumRealisedPnl()

// 	require.NoError(s.T(), err)
// 	require.Nil(s.T(), deep.Equal(TestSum, sum))

// }
