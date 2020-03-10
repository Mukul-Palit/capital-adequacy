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

type PositionSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	pos               Pos
	position_snapshot *PositionSnapshot
}

func (p *PositionSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, p.mock, err = sqlmock.New()
	require.NoError(p.T(), err)

	p.DB, err = gorm.Open("mysql", db)
	require.NoError(p.T(), err)

	//s.DB.LogMode(true)

	p.pos = CreatePositionSnapshot(p.DB)
}

func (p *PositionSuite) AfterTest(_, _ string) {
	require.NoError(p.T(), p.mock.ExpectationsWereMet())
}

func TestPosInit(t *testing.T) {
	suite.Run(t, new(PositionSuite))
}

func (p *PositionSuite) TestGetPositionSnapshot() {
	var (
		id        = "1"
		timestamp = "2020-02-24 10:10:10"
		alpha     = "0.99"
		days      = "1"
		label     = "Warehouse-2011Parametric"
		exposure  = 100.12
		Var       = -100.5666
	)
	selectQuery := fmt.Sprintf("SELECT id,ts,alpha,days,label,exposure,var FROM `var`  WHERE \\(alpha \\= 0\\.99 AND days \\= 1 AND label \\= 'Warehouse\\-2011Parametric'\\) ORDER BY id desc LIMIT 1")
	p.mock.ExpectQuery(selectQuery).
		WillReturnRows(sqlmock.NewRows([]string{"id", "timestamp", "alpha", "days", "label", "exposure", "Var"}).
			AddRow(id, timestamp, alpha, days, label, exposure, Var))
	res, err := p.pos.GetPositionSnapshot()
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(p.T(), err)
	require.Nil(p.T(), deep.Equal(PositionSnapshot{ID: id, Timestamp: timestamp, Alpha: alpha, Days: days, Label: label, Exposure: exposure, Var: Var}, res))
}
